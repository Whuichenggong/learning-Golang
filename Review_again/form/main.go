package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type LoginData struct {
	Token   string
	Message string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		// 通过MD5(时间戳)来获取唯一值，然后我们把这个值存储到服务器端 session来控制
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timestamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))

		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Println("Error parsing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// 将 token 传递给模板
		data := LoginData{Token: token}
		err = t.Execute(w, data)
		if err != nil {
			log.Println("Error rendering template:", err)
		}

		//t.Execute(w, data) // 传递 nil 表示没有动态消息

	} else {

		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token == "" {
			//验证token的合法性
			fmt.Println("Invalid token")
		} else {
			// 验证 token 的合法性
			fmt.Println("Token received:", token)
		}

		fmt.Println("fruit:", r.Form["fruit"])
		fmt.Println("gender:", r.Form["gender"])
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端

		// 获取表单字段
		username := r.Form.Get("username")
		fruit := r.Form.Get("fruit")
		gender := r.Form.Get("gender")

		// 进行字段验证 用map比下面对比简单
		validFruits := map[string]bool{"apple": true, "banana": true, "orange": true}
		if _, valid := validFruits[fruit]; valid {
			fmt.Println("Valid fruit:", fruit)
		} else {
			fmt.Println("Invalid fruit:", fruit)
		}

		validGenders := map[string]bool{"1": true, "2": true}
		if _, valid := validGenders[gender]; valid {
			fmt.Println("Valid gender:", gender)
		} else {
			fmt.Println("Invalid gender:", gender)
		}

		//这段代码太冗余了，但是可以作为一种方法
		// slice := []string{"apple", "banana", "orange"}
		// //判断水果是否在可选列表中判断这个值是否是我们预设的值
		// v := r.Form.Get("fruit")
		// for _, item := range slice {
		// 	if item == v {
		// 		fmt.Println("fruit is valid")
		// 	}
		// }
		// slice1 := []string{"1", "2"}

		// for _, v := range slice1 {
		// 	if v == r.Form.Get("gender") {
		// 		fmt.Println("gender is valid")
		// 	}
		// }

		// 验证逻辑（示例）
		message := fmt.Sprintf("Hello %s, your favorite fruit is %s.", username, fruit)
		// 加载模板并渲染页面（带动态消息）
		data := LoginData{Token: token, Message: message}
		t, err := template.ParseFiles("login.gtpl")
		if err != nil {
			log.Println("Error parsing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// 渲染模板，并传递动态数据
		err = t.Execute(w, data)
		if err != nil {
			log.Println("Error rendering template:", err)
		}
	}

}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
