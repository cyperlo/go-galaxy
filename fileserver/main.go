package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前运行目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前目录失败: %v", err)
	}

	// 打印提示信息
	fmt.Printf("Serving files from: %s\n", dir)
	fmt.Println("访问 http://localhost:8080 即可浏览当前目录下的文件")

	resourcesDir := filepath.Join(dir, "resources")
	// 使用 http.FileServer 处理静态文件请求
	fs := http.FileServer(http.Dir(resourcesDir))

	// 路由设置，将根路径映射到静态文件服务器
	http.Handle("/", fs)

	// 启动服务器
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动 HTTP 服务器失败: %v", err)
	}
}
