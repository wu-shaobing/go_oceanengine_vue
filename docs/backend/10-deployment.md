# 部署方案

## 概述

本文档描述系统的部署架构、容器化配置、CI/CD 流程和运维监控方案。

## 部署架构

### 生产环境架构

```
                                    ┌─────────────────┐
                                    │   CDN / OSS     │
                                    └────────┬────────┘
                                             │
┌──────────────────────────────────────────────────────────────────────┐
│                           Kubernetes Cluster                          │
│                                                                        │
│  ┌─────────────┐    ┌─────────────────────────────────────────────┐  │
│  │   Ingress   │────│                  Service                    │  │
│  │  (Nginx)    │    └─────────────────────────────────────────────┘  │
│  └─────────────┘                        │                            │
│                                         │                            │
│  ┌──────────────────────────────────────┼─────────────────────────┐  │
│  │                                      │                         │  │
│  │  ┌───────────┐  ┌───────────┐  ┌───────────┐                   │  │
│  │  │  API Pod  │  │  API Pod  │  │  API Pod  │  <- Deployment    │  │
│  │  │  (Go)     │  │  (Go)     │  │  (Go)     │                   │  │
│  │  └───────────┘  └───────────┘  └───────────┘                   │  │
│  │                                                                │  │
│  │  ┌───────────┐  ┌───────────┐                                  │  │
│  │  │  Worker   │  │  Worker   │  <- Deployment (定时任务/异步)    │  │
│  │  │  Pod      │  │  Pod      │                                  │  │
│  │  └───────────┘  └───────────┘                                  │  │
│  │                                                                │  │
│  └────────────────────────────────────────────────────────────────┘  │
│                                                                        │
└──────────────────────────────────────────────────────────────────────┘
         │                    │                    │
         ▼                    ▼                    ▼
┌─────────────┐      ┌─────────────┐      ┌─────────────┐
│   MySQL     │      │   Redis     │      │ Prometheus  │
│   Cluster   │      │   Cluster   │      │  + Grafana  │
└─────────────┘      └─────────────┘      └─────────────┘
```

---

## Docker 配置

### Dockerfile

```dockerfile
# Dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git ca-certificates tzdata

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源码
COPY . .

# 构建
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X main.Version=${VERSION:-dev} -X main.BuildTime=$(date -u +%Y%m%d%H%M%S)" \
    -o /app/server ./cmd/server

# 运行阶段
FROM alpine:3.18

WORKDIR /app

# 安装必要的工具
RUN apk add --no-cache ca-certificates tzdata curl

# 设置时区
ENV TZ=Asia/Shanghai

# 创建非 root 用户
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

# 复制二进制文件
COPY --from=builder /app/server /app/server
COPY --from=builder /app/config /app/config

# 设置权限
RUN chown -R appuser:appgroup /app

# 切换用户
USER appuser

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# 暴露端口
EXPOSE 8080

# 启动命令
ENTRYPOINT ["/app/server"]
CMD ["-c", "/app/config/config.yaml"]
```

### Docker Compose (开发环境)

```yaml
# docker-compose.yml
version: '3.8'

services:
  api:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - CONFIG_PATH=/app/config/config.yaml
      - GIN_MODE=debug
    volumes:
      - ./config:/app/config:ro
      - ./logs:/app/logs
    depends_on:
      mysql:
        condition: service_healthy
      redis:
        condition: service_healthy
    networks:
      - oceanengine
    restart: unless-stopped

  worker:
    build:
      context: .
      dockerfile: Dockerfile.worker
    environment:
      - CONFIG_PATH=/app/config/config.yaml
    volumes:
      - ./config:/app/config:ro
      - ./logs:/app/logs
    depends_on:
      - mysql
      - redis
    networks:
      - oceanengine
    restart: unless-stopped

  mysql:
    image: mysql:8.0
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: root123
      MYSQL_DATABASE: oceanengine
      MYSQL_USER: oceanengine
      MYSQL_PASSWORD: oceanengine123
    volumes:
      - mysql_data:/var/lib/mysql
      - ./scripts/sql:/docker-entrypoint-initdb.d
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - oceanengine

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - oceanengine

  prometheus:
    image: prom/prometheus:v2.45.0
    ports:
      - "9090:9090"
    volumes:
      - ./deploy/prometheus:/etc/prometheus
      - prometheus_data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
    networks:
      - oceanengine

  grafana:
    image: grafana/grafana:10.0.0
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
    volumes:
      - grafana_data:/var/lib/grafana
      - ./deploy/grafana/dashboards:/etc/grafana/provisioning/dashboards
      - ./deploy/grafana/datasources:/etc/grafana/provisioning/datasources
    networks:
      - oceanengine

volumes:
  mysql_data:
  redis_data:
  prometheus_data:
  grafana_data:

networks:
  oceanengine:
    driver: bridge
```

---

## Kubernetes 配置

### Namespace

```yaml
# k8s/namespace.yaml
apiVersion: v1
kind: Namespace
metadata:
  name: oceanengine
  labels:
    name: oceanengine
```

### ConfigMap

```yaml
# k8s/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: oceanengine-config
  namespace: oceanengine
data:
  config.yaml: |
    server:
      mode: release
      port: 8080
      read_timeout: 30s
      write_timeout: 30s
    
    database:
      driver: mysql
      host: mysql-service
      port: 3306
      database: oceanengine
      charset: utf8mb4
      max_idle_conns: 10
      max_open_conns: 100
      conn_max_lifetime: 1h
    
    redis:
      addr: redis-service:6379
      db: 0
      pool_size: 100
    
    logger:
      level: info
      format: json
      output: stdout
    
    oceanengine:
      base_url: https://ad.oceanengine.com/open_api
      timeout: 30s
```

### Secret

```yaml
# k8s/secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: oceanengine-secret
  namespace: oceanengine
type: Opaque
stringData:
  DB_PASSWORD: "your-db-password"
  REDIS_PASSWORD: "your-redis-password"
  JWT_SECRET: "your-jwt-secret"
  OE_APP_ID: "your-oceanengine-app-id"
  OE_SECRET: "your-oceanengine-secret"
```

### Deployment

```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: oceanengine-api
  namespace: oceanengine
  labels:
    app: oceanengine-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: oceanengine-api
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  template:
    metadata:
      labels:
        app: oceanengine-api
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
        prometheus.io/path: "/metrics"
    spec:
      containers:
        - name: api
          image: oceanengine/api:latest
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
              name: http
          envFrom:
            - secretRef:
                name: oceanengine-secret
          volumeMounts:
            - name: config
              mountPath: /app/config
              readOnly: true
          resources:
            requests:
              cpu: "100m"
              memory: "128Mi"
            limits:
              cpu: "500m"
              memory: "512Mi"
          livenessProbe:
            httpGet:
              path: /health
              port: 8080
            initialDelaySeconds: 10
            periodSeconds: 10
            timeoutSeconds: 5
            failureThreshold: 3
          readinessProbe:
            httpGet:
              path: /health
              port: 8080
            initialDelaySeconds: 5
            periodSeconds: 5
            timeoutSeconds: 3
            failureThreshold: 3
      volumes:
        - name: config
          configMap:
            name: oceanengine-config
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
            - weight: 100
              podAffinityTerm:
                labelSelector:
                  matchLabels:
                    app: oceanengine-api
                topologyKey: kubernetes.io/hostname
```

### Service

```yaml
# k8s/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: oceanengine-api
  namespace: oceanengine
  labels:
    app: oceanengine-api
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
      name: http
  selector:
    app: oceanengine-api
```

### Ingress

```yaml
# k8s/ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: oceanengine-ingress
  namespace: oceanengine
  annotations:
    nginx.ingress.kubernetes.io/proxy-body-size: "50m"
    nginx.ingress.kubernetes.io/proxy-read-timeout: "60"
    nginx.ingress.kubernetes.io/proxy-send-timeout: "60"
    cert-manager.io/cluster-issuer: "letsencrypt-prod"
spec:
  ingressClassName: nginx
  tls:
    - hosts:
        - api.oceanengine.example.com
      secretName: oceanengine-tls
  rules:
    - host: api.oceanengine.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: oceanengine-api
                port:
                  number: 80
```

### HPA (自动扩缩容)

```yaml
# k8s/hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: oceanengine-api-hpa
  namespace: oceanengine
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: oceanengine-api
  minReplicas: 3
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70
    - type: Resource
      resource:
        name: memory
        target:
          type: Utilization
          averageUtilization: 80
```

---

## CI/CD 配置

### GitHub Actions

```yaml
# .github/workflows/deploy.yaml
name: Deploy

on:
  push:
    branches: [main]
    tags: ['v*']

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out

  build:
    needs: test
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write
    outputs:
      version: ${{ steps.meta.outputs.version }}
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      
      - name: Log in to Container Registry
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      
      - name: Extract metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
          tags: |
            type=ref,event=branch
            type=semver,pattern={{version}}
            type=sha
      
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max

  deploy-staging:
    needs: build
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    environment: staging
    steps:
      - uses: actions/checkout@v4
      
      - name: Deploy to staging
        uses: azure/k8s-deploy@v4
        with:
          namespace: oceanengine-staging
          manifests: |
            k8s/deployment.yaml
            k8s/service.yaml
          images: |
            ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:${{ needs.build.outputs.version }}

  deploy-production:
    needs: build
    if: startsWith(github.ref, 'refs/tags/v')
    runs-on: ubuntu-latest
    environment: production
    steps:
      - uses: actions/checkout@v4
      
      - name: Deploy to production
        uses: azure/k8s-deploy@v4
        with:
          namespace: oceanengine
          manifests: |
            k8s/deployment.yaml
            k8s/service.yaml
          images: |
            ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:${{ needs.build.outputs.version }}
```

---

## 监控配置

### Prometheus 配置

```yaml
# deploy/prometheus/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

alerting:
  alertmanagers:
    - static_configs:
        - targets: ['alertmanager:9093']

rule_files:
  - "/etc/prometheus/rules/*.yml"

scrape_configs:
  - job_name: 'oceanengine-api'
    kubernetes_sd_configs:
      - role: pod
        namespaces:
          names: ['oceanengine']
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
        action: keep
        regex: true
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
        action: replace
        target_label: __metrics_path__
        regex: (.+)
```

### 告警规则

```yaml
# deploy/prometheus/rules/alerts.yml
groups:
  - name: oceanengine
    rules:
      - alert: HighErrorRate
        expr: sum(rate(http_requests_total{status=~"5.."}[5m])) / sum(rate(http_requests_total[5m])) > 0.05
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "High error rate detected"
          description: "Error rate is {{ $value | humanizePercentage }}"

      - alert: HighLatency
        expr: histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) > 1
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High latency detected"
          description: "95th percentile latency is {{ $value }}s"

      - alert: PodRestart
        expr: increase(kube_pod_container_status_restarts_total{namespace="oceanengine"}[1h]) > 3
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "Pod restart detected"
          description: "Pod {{ $labels.pod }} has restarted {{ $value }} times"
```

---

## 运维脚本

### 数据库备份

```bash
#!/bin/bash
# scripts/backup-db.sh

DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/data/backups/mysql"
RETENTION_DAYS=30

# 创建备份目录
mkdir -p $BACKUP_DIR

# 执行备份
mysqldump -h $DB_HOST -u $DB_USER -p$DB_PASSWORD \
    --single-transaction \
    --routines \
    --triggers \
    oceanengine > $BACKUP_DIR/oceanengine_$DATE.sql

# 压缩备份
gzip $BACKUP_DIR/oceanengine_$DATE.sql

# 上传到 OSS
aliyun oss cp $BACKUP_DIR/oceanengine_$DATE.sql.gz \
    oss://oceanengine-backup/mysql/

# 清理旧备份
find $BACKUP_DIR -name "*.sql.gz" -mtime +$RETENTION_DAYS -delete
```

### 健康检查

```bash
#!/bin/bash
# scripts/health-check.sh

API_URL="http://localhost:8080/health"
MAX_RETRIES=3
RETRY_INTERVAL=5

check_health() {
    response=$(curl -s -o /dev/null -w "%{http_code}" $API_URL)
    if [ "$response" == "200" ]; then
        return 0
    fi
    return 1
}

retry=0
while [ $retry -lt $MAX_RETRIES ]; do
    if check_health; then
        echo "Health check passed"
        exit 0
    fi
    
    retry=$((retry + 1))
    echo "Health check failed, retry $retry/$MAX_RETRIES"
    sleep $RETRY_INTERVAL
done

echo "Health check failed after $MAX_RETRIES retries"
exit 1
```

---

## 环境配置

### 开发环境

```yaml
# config/config.dev.yaml
server:
  mode: debug
  port: 8080

database:
  host: localhost
  port: 3306
  database: oceanengine_dev

redis:
  addr: localhost:6379
  db: 0

logger:
  level: debug
  format: console
  output: stdout
```

### 生产环境

```yaml
# config/config.prod.yaml
server:
  mode: release
  port: 8080
  read_timeout: 30s
  write_timeout: 30s

database:
  host: ${DB_HOST}
  port: 3306
  database: oceanengine
  username: ${DB_USER}
  password: ${DB_PASSWORD}
  max_idle_conns: 10
  max_open_conns: 100

redis:
  addr: ${REDIS_ADDR}
  password: ${REDIS_PASSWORD}
  db: 0
  pool_size: 100

logger:
  level: info
  format: json
  output: file
  filename: /var/log/oceanengine/app.log

jwt:
  secret: ${JWT_SECRET}
  access_expire: 2h
  refresh_expire: 168h

oceanengine:
  app_id: ${OE_APP_ID}
  secret: ${OE_SECRET}
```
