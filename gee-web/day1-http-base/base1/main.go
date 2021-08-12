package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)
	// 下面一行用来启动web服务。第二个参数就是我们基于net/http标准库实现web框架的入口
	log.Fatal(http.ListenAndServe(":9999",nil))
	/* http.Handler接口定义如下
	只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了
	具体实现见base2/main.go
	package http

	type Handler interface {
		ServeHTTP(w ResponseWriter, r *Request)
	}

	func ListenAndServe(address string, h Handler) error
	 */
}

func indexHandler(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "URL = %q\n", req.URL)
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req * http.Request){
	for k, v := range req.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n",k,v)
	}
}