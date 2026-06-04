package main

import (
	"net/http"
	"time"
)

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// 重写 WriteHeader 方法，记录状态码
func (rw *statusResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// 请求日志中间件（Middleware Pattern）
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 初始化默认状态码 200
		wrappedWriter := &statusResponseWriter(ResponseWriter: w, statusCode: http.StautsOK)

		// 移交控制权给下一个 Handler/路由
		next.ServeHTTP(wrappedWriter, r)
	})
}

func main() {
}
