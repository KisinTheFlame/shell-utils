package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse 定义健康检查响应结构
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	// 注册健康检查处理函数
	http.HandleFunc("/health", healthHandler)

	// 在32001端口启动服务器
	port := ":32001"
	println("服务器启动，监听端口", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		println("服务器启动失败:", err.Error())
	}
}

// healthHandler 处理健康检查请求
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为JSON类型
	w.Header().Set("Content-Type", "application/json")

	// 创建健康响应数据
	response := HealthResponse{
		Status:    "on",
		Timestamp: time.Now(),
	}

	// 将响应数据序列化为JSON并返回
	json.NewEncoder(w).Encode(response)
}
