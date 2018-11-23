package main

import (
	"net/http"
	"fmt"
)

type MyMux struct {

}

/**
type Handler interface {
ServeHTTP(ResponseWriter, *Request) //  路由实现器
}
 */

//这里是实现接口
func(p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.URL.Path == "/" {
		sayHelloName2(w,r)
		return
	}
	http.NotFound(w,r)
	return
}
func sayHelloName2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer,"hello,myroute")
}


func main() {
	mux:=&MyMux{}
	http.ListenAndServe("127.0.0.1:9090",mux)

}

/**
Go代码的执行流程
通过对http包的分析之后，现在让我们来梳理一下整个的代码执行过程。
首先调用Http.HandleFunc
按顺序做了几件事：
1 调用了DefaultServerMux的HandleFunc
2 调用了DefaultServerMux的Handle
3 往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则
其次调用http.ListenAndServe(":9090", nil)
按顺序做了几件事情：
1 实例化Server
2 调用Server的ListenAndServe()
3 调用net.Listen("tcp", addr)监听端口
4 启动一个for循环，在循环体中Accept请求
5 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()
6 读取每个请求的内容w, err := c.readRequest()
7 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为
DefaultServeMux
8 调用handler的ServeHttp
9 在这个例子中，下面就进入到DefaultServerMux.ServeHttp
10 根据request选择handler，并且进入到这个handler的ServeHTTP
mux.handler(r).ServeHTTP(w, r)

11 选择handler：
A 判断是否有路由能满足这个request（循环遍历ServerMux的muxEntry）
B 如果有路由满足，调用这个路由handler的ServeHttp
C 如果没有路由满足，调用NotFoundHandler的ServeHttp
 */
