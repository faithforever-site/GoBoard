package handlers

import (
	"GoBoard/tools"
	"fmt"
	"html/template"
	"net/http"
)

// HomeHandler 渲染后台首页
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL.Path, "Method:", r.Method, "Content-Type:", r.Header.Get("Content-Type"))
	params := tools.GetAllParams(r)
	fmt.Println("返回结果：", params)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "模板解析错误", http.StatusInternalServerError)
		return
	}
	_ = tmpl.Execute(w, map[string]any{
		"Title":  "后台首页",
		"Params": params,
	})
}
