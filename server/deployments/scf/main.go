// 腾讯云 SCF 入口适配器
// 使用 API 网关触发器将 HTTP 请求转发到 Gin
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/bububa/oceanengine/server/internal/config"
	"github.com/bububa/oceanengine/server/internal/handler"
	"github.com/bububa/oceanengine/server/internal/middleware"
	"github.com/bububa/oceanengine/server/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

// APIGatewayRequest API网关请求结构
type APIGatewayRequest struct {
	Path              string            `json:"path"`
	HTTPMethod        string            `json:"httpMethod"`
	Headers           map[string]string `json:"headers"`
	QueryString       map[string]string `json:"queryString"`
	Body              string            `json:"body"`
	IsBase64Encoded   bool              `json:"isBase64Encoded"`
	PathParameters    map[string]string `json:"pathParameters"`
	StageVariables    map[string]string `json:"stageVariables"`
	RequestContext    map[string]any    `json:"requestContext"`
}

// APIGatewayResponse API网关响应结构
type APIGatewayResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

var router *gin.Engine

func init() {
	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)

	// 加载配置
	cfg := config.Load()

	// 创建服务
	oceanEngineSvc := service.NewOceanEngineService(
		cfg.OceanEngine.AppID,
		cfg.OceanEngine.AppSecret,
	)

	// 创建Handler
	h := handler.NewHandler(oceanEngineSvc)

	// 创建路由
	router = gin.New()
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	setupRoutes(router, h)
}

func setupRoutes(r *gin.Engine, h *handler.Handler) {
	r.GET("/health", h.Health)

	v1 := r.Group("/api/v1")
	{
		oauth := v1.Group("/oauth")
		{
			oauth.POST("/auth_url", h.GetAuthURL)
			oauth.POST("/access_token", h.GetAccessToken)
			oauth.POST("/refresh_token", h.RefreshToken)
		}

		advertiser := v1.Group("/advertiser")
		{
			advertiser.GET("/list", h.GetAdvertisers)
			advertiser.POST("/info", h.GetAdvertiserInfo)
		}
	}
}

// Handler SCF 入口函数
func Handler(ctx context.Context, event APIGatewayRequest) (APIGatewayResponse, error) {
	// 创建HTTP请求
	req, err := http.NewRequest(event.HTTPMethod, event.Path, strings.NewReader(event.Body))
	if err != nil {
		return APIGatewayResponse{
			StatusCode: 500,
			Body:       `{"error": "Failed to create request"}`,
		}, nil
	}

	// 设置headers
	for k, v := range event.Headers {
		req.Header.Set(k, v)
	}

	// 设置query参数
	q := req.URL.Query()
	for k, v := range event.QueryString {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 处理请求
	router.ServeHTTP(w, req)

	// 构建响应headers
	respHeaders := make(map[string]string)
	for k, v := range w.Header() {
		if len(v) > 0 {
			respHeaders[k] = v[0]
		}
	}

	return APIGatewayResponse{
		StatusCode: w.Code,
		Headers:    respHeaders,
		Body:       w.Body.String(),
	}, nil
}

func main() {
	// 判断是否在云函数环境运行
	if os.Getenv("SCF_RUNTIME") != "" || os.Getenv("_SCF_SERVER_PORT") != "" {
		cloudfunction.Start(Handler)
	} else {
		// 本地调试模式
		testEvent := APIGatewayRequest{
			Path:       "/health",
			HTTPMethod: "GET",
		}
		resp, _ := Handler(context.Background(), testEvent)
		output, _ := json.MarshalIndent(resp, "", "  ")
		println(string(output))
	}
}
