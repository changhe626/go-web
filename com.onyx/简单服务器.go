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



如何监听端口？
如何接收客户端请求？
如何分配handler？


1.Go是通过一个函数 ListenAndServe 来处理这些事情的，这个底层其实这样处
理的：初始化一个server对象，然后调用了 net.Listen("tcp", addr) ，也就是底层用TCP协议搭建了一个服
务，然后监控我们设置的端口。

2.执行监控端口之后，调用了 srv.Serve(net.Listener) 函数，这个
函数就是处理接收客户端的请求信息。这个函数里面起了一个 for{} ，首先通过Listener接收请求，其次创建一个
Conn，最后单独开了一个goroutine，把这个请求的数据当做参数扔给这个conn去服务： go c.serve() 。这个就是
高并发体现了，用户的每一次请求都是在一个新的goroutine去服务，相互不影响。


3.conn首先会解析request: c.readRequest() ,然后获取相应的
handler: handler := c.server.Handler ，也就是我们刚才在调用函数 ListenAndServe 时候的第二个参数，
我们前面例子传递的是nil，也就是为空，那么默认获取 handler = DefaultServeMux ,那么这个变量用来做什么
的呢？对，这个变量就是一个路由器，它用来匹配url跳转到其相应的handle函数，那么这个我们有设置过吗?有，我
们调用的代码里面第一句不是调用了 http.HandleFunc("/", sayhelloName) 嘛。这个作用就是注册了请求 / 的
路由规则，当请求uri为"/"，路由就会转到函数sayhelloName，DefaultServeMux会调用ServeHTTP方法，这个方法内
部其实就是调用sayhelloName本身，最后通过写入response的信息反馈到客户端。








Go为了实现高并发和高性能, 使用了goroutines来处理Conn的读写事件, 这样每
个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是Go高效的保证。

 */
