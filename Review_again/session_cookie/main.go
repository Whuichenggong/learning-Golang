package main

import (
	"log"
	"net/http"
	"session/session"
	"text/template"
)

var globalSessions *session.Manager

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

// 然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
}

func main() {
	// 这里添加 HTTP 路由和请求处理
	http.HandleFunc("/login", login)

	// 启动 HTTP 服务器
	log.Println("Starting server on :8888...")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
