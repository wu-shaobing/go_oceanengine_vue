package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	// 简单的健康检查端点
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "ok", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
	})

	// 简单的API信息端点
	http.HandleFunc("/api/v1/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"app": "OceanEngine Backend",
			"version": "1.0.0",
			"environment": "CloudStudio",
			"timestamp": "%s"
		}`, time.Now().Format(time.RFC3339))
	})

	// 简单的广告主列表端点（模拟数据）
	http.HandleFunc("/api/v1/advertisers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"code": 0,
			"message": "success",
			"data": {
				"list": [
					{
						"id": 1,
						"name": "测试广告主1",
						"status": "active",
						"created_at": "2023-01-01T00:00:00Z"
					},
					{
						"id": 2,
						"name": "测试广告主2",
						"status": "active",
						"created_at": "2023-01-02T00:00:00Z"
					}
				],
				"total": 2,
				"page": 1,
				"page_size": 10
			}
		}`)
	})

	// 处理前端请求的所有其他路径，返回简单的响应
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Printf("请求路径: %s\n", path)

		// 对于API请求，返回JSON响应
		if strings.HasPrefix(path, "/api/") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{
				"code": 0,
				"message": "OceanEngine Backend API - CloudStudio版本",
				"path": "%s",
				"timestamp": "%s"
			}`, path, time.Now().Format(time.RFC3339))
			return
		}

		// 对于非API请求，返回HTML响应
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>OceanEngine Backend - CloudStudio</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        .api-list { list-style-type: none; padding: 0; }
        .api-item { margin: 10px 0; padding: 10px; background: #f5f5f5; border-radius: 5px; }
        code { background: #e8e8e8; padding: 2px 4px; border-radius: 3px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🌊 OceanEngine Backend API</h1>
        <p>CloudStudio 部署版本 - 运行正常</p>
        
        <h2>可用的 API 端点：</h2>
        <ul class="api-list">
            <li class="api-item">
                <strong>健康检查</strong><br>
                <code>GET /health</code>
            </li>
            <li class="api-item">
                <strong>API 信息</strong><br>
                <code>GET /api/v1/info</code>
            </li>
            <li class="api-item">
                <strong>广告主列表</strong><br>
                <code>GET /api/v1/advertisers</code>
            </li>
        </ul>
        
        <p><strong>当前时间：</strong> %s</p>
    </div>
</body>
</html>`, time.Now().Format(time.RFC3339))
	})

	fmt.Printf("🚀 OceanEngine Backend (CloudStudio版本) 启动成功！\n")
	fmt.Printf("📍 服务地址: http://localhost:%s\n", port)
	fmt.Printf("🏥 健康检查: http://localhost:%s/health\n", port)
	fmt.Printf("📊 API 信息: http://localhost:%s/api/v1/info\n", port)
	fmt.Printf("📱 广告主列表: http://localhost:%s/api/v1/advertisers\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
		os.Exit(1)
	}
}
