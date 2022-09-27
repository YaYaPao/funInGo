package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Creating a data structure with load and save methods
// Using the net/http package to build web applications
// Using the html/template package to process HTML templates
// Using the regexp package to validate user input
// Using closures

type Page struct {
	Title string
	Body  []byte
}

// 读取文件
func LoadFile(t string) (*Page, error) {
	filename := t + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: filename, Body: body}, nil
}

// 每个路由处理器包含两个传参：
//  - `http.ResponseWriter` 用来写入返回数据
//  - `*http.Request` 用来获取 request
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi, path is %s!", r.URL.Path[1:])
	if err != nil {
		return
	}
}

// 读取指定文件
// 注意引入 wiki 包的方式
func viewWikiHandler(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Path[len("/view/"):]
	p, err := LoadFile("./wiki/" + filepath)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Fprintf(w, "<h1>%s</h1><main>%s</main>", p.Title[len("./wiki/"):], p.Body)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	_ = t.Execute(w, p)
}

// viewHandler 当访问 `http://localhost:8088/view/hello`, 时：
//  - 从 hello.txt 内读取文件内容
//  - 将 Page 内容写入 view.html 的模版内
//  - 向浏览器输出模版内容
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := LoadFile(title)
	if err != nil {
		fmt.Println(err)
	}
	renderTemplate(w, "view", p)
}

// 通过 `http.HandleFunc` 将路由绑定到指定处理器
// 通过 `http.ListenAndServer` 来监听指定端口，通过 `log.Fatal` 来输出错误
func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
