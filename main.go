package main

import (
	"fmt"
	"log"
	"net/http"

	"GoBoard/db"
	"GoBoard/handlers"
)

// 打印 ASCII Logo
func printLogo() {
	logo := `
   ____       ____               _ 
  / ___| ___ | __ )  __ _ _ __  | |
 | |  _ / _ \|  _ \ / _' | '_ \ | |
 | |_| | (_) | |_) | (_| | | | ||_|
  \____|\___/|____/ \__,_|_| |_|(_)
  
      🚀 Welcome to GoBoard 🚀
`
	fmt.Println(logo)
}

func main() {
	// 打印 Logo
	printLogo()

	// 初始化数据库
	if err := db.InitDB("goboard.db"); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	fmt.Println("✅ 数据库连接成功")

	// 路由
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/api/hello", handlers.HelloHandler)

	// 启动服务
	addr := ":8080"
	fmt.Printf("🚀 GoBoard 已启动，访问 http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
