package main

// go语言官方之外的包是引用gopath的
// 1.6后提供了vendor 把所有的项目依赖独立起来 一个项目一个vendor
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	//"github.com/thinkeridea/go-extend/tree/master/exnet"
)

func Healthz(ctx *Context) {
	//4.当访问 localhost/healthz 时，应返回200
	io.WriteString(ctx.W, "200")
	ctx.W.WriteHeader(200)
}

func Task1(ctx *Context) {
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	fmt.Fprintf(ctx.W, "client request header:\n")
	count := 0
	for k, v := range ctx.R.Header {
		ctx.W.Header().Set(k, v[0])
		fmt.Fprintf(ctx.W, "%d : %s -- > %s\n", count, k, v)
		count++
	}
}

func Task2(ctx *Context) {
	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	environ := os.Environ()

	for index, s := range environ {
		ctx.W.Header().Set(string(index), s)
	}
	fmt.Fprintf(ctx.W, "your head has changed")
}

func Task3(ctx *Context) {
	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

	ctx.W.WriteHeader(http.StatusOK)
	fmt.Print(time.Now())
	fmt.Printf("--The request HTTP code is %d, the client IP is %s\n", http.StatusOK, ctx.R.RemoteAddr)

}

func main() {
	server := NewHttpServer("cncamp-httpserver")
	server.Route("/task1/", Task1)
	server.Route("/task2/", Task2)
	server.Route("/task3/", Task3)
	server.Route("/healthz/", Healthz)
	server.Start(":8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
