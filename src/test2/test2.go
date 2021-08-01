package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("../bouncing-balls-start/index.html", "../bouncing-balls-start/main.js", "../bouncing-balls-start/style.css")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", sayHello)
	//静态文件处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../bouncing-balls-start"))))
	//html文件处理

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}

}
