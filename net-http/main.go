package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ===== 数据模型 ======
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 模拟数据库
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}
var nextID = 3

// ====== 中间件登录 ======
// 这是一个典型的中间件模式: 接收一个 Handler， 返回一个新的 Handler
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// 调用真正的处理器
		next.ServeHTTP(w, r)
		// 记录日志
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// ======= 路由分发器 =========
// 解析URL 路径，把请求发给对应的 handler
func router(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")

	// 用户相关的内容
	if strings.HasPrefix(path, "/users") {
		usersHandler(w, r, path)
		return
	}
	http.NotFound(w, r)
}

func usersHandler(w http.ResponseWriter, r *http.Request, path string) {
	// GET /api/users -> 获取用户列表
	// GET /api/users/1 -> 获取单个用户
	// POST /api/users -> 创建用户

	switch r.Method {
	case http.MethodGet:
		if path == "/users" {
			listUsers(w, r)
		} else {
			// 提取路径中的 ID: /users/1 -> id=1
			idStr := strings.TrimPrefix(path, "/users/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "无效的用户ID", http.StatusBadRequest)
				return
			}
			getUser(w, r, id)
		}

	case http.MethodPost:
		if path == "/users" {
			createUser(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// GET: /api/users
func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// POST: /api/users
func createUser(w http.ResponseWriter, r *http.Request) {
	// 解析请求体
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "格式请求错误:", http.StatusBadRequest)
		return
	}
	// 分配ID到数据库
	newUser.ID = nextID
	nextID++
	// 4. 返回201 created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func getUser(w http.ResponseWriter, r *http.Request, id int) {
	for _, u := range users {
		if u.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "用户不存在", http.StatusNotFound)
}

// 程序入口

func main() {
	// 创建路由器
	mux := http.NewServeMux()

	// 注册路由： 所有/api/ 开头的请求交给router 处理
	mux.HandleFunc("/api/", router)

	// 健康检查接口（不需要 /api 前缀）
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	// 包装中间件
	handler := loggingMiddleware(mux)

	// 启动服务器
	server := &http.Server{
		Addr:         ":9090",
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Println("服务器情动在 http://localhost:9090")
	log.Println("可用接口:")
	log.Println("  GET  /health")
	log.Println("  GET  /api/users")
	log.Println("  GET  /api/users/1")
	log.Println("  POST /api/users")
	// ListenAndServe 会阻塞  知道出错才返回
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
