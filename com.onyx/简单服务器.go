package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm() // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的

}


func main() {
	http.HandleFunc("/test", sayHelloName) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

/**
web工作方式的几个概念
以下均是服务器端的几个概念
Request：用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息
Response：服务器需要反馈给客户端的信息
Conn：用户的每次请求链接
Handler：处理请求和生成返回信息的处理逻辑


http包执行流程
1. 创建Listen Socket, 监听指定的端口, 等待客户端请求到来。
2. Listen Socket接受客户端的请求, 得到Client Socket, 接下来通过Client Socket与客户端通信。
3. 处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头, 如果是POST方法, 还可能要读取客户端提
交的数据, 然后交给相应的handler处理请求, handler处理完毕准备好客户端需要的数据, 通过Client
Socket写给客户端。


Go为了实现高并发和高性能, 使用了goroutines来处理Conn的读写事件, 这样每
个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是Go高效的保证。
 */
