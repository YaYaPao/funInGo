package main

import (
	"fmt"
	"log"
	"net/http"
	"wiki"
)

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
func viewHandler(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Path[len("/view/"):]
	p, err := wiki.LoadFile("./wiki/" + filepath)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Fprintf(w, "<h1>%s</h1><main>%s</main>", p.Title[len("./wiki/"):], p.Body)
}

// 通过 `http.HandleFunc` 将路由绑定到指定处理器
// 通过 `http.ListenAndServer` 来监听指定端口，通过 `log.Fatal` 来输出错误
func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
