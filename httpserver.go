package main

// go语言官方之外的包是引用gopath的
// 1.6后提供了vendor 把所有的项目依赖独立起来 一个项目一个vendor
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	//"github.com/thinkeridea/go-extend/tree/master/exnet"
)

func task1(w http.ResponseWriter, r *http.Request) {
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	fmt.Fprintf(w, "client request header:\n")
	count := 0
	for k, v := range r.Header {
		fmt.Fprintf(w, "%d : %s -- > %s\n", count, k, v)
		count++
	}
}

func task2(w http.ResponseWriter, r *http.Request) {
	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	environ := os.Environ()

	for index, s := range environ {
		w.Header().Set(string(index), s)
	}

}

//func RemoteIp(req *http.Request) string {
//	remoteAddr := req.RemoteAddr
//	if ip := req.Header.Get(XRealIP); ip != "" {
//		remoteAddr = ip
//	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
//		remoteAddr = ip
//	} else {
//		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
//	}
//
//	if remoteAddr == "::1" {
//		remoteAddr = "127.0.0.1"
//	}
//
//	return remoteAddr
//}

func task3(w http.ResponseWriter, r *http.Request) {
	//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

	w.WriteHeader(http.StatusOK)
	fmt.Print(time.Now())
	fmt.Printf("--The request HTTP code is %d, the client IP is %s\n", http.StatusOK, r.RemoteAddr)
	//fmt.Println(http.StatusOK)
	//fmt.Println(r.RemoteAddr)
	//fmt.Println(r.Header.Get("Remote_addr"))

	//fileName := "ll.log"
	//logFile,err  := os.Create(fileName)
	//defer logFile.Close()
	//if err != nil {
	//	log.Fatalln("open file error !")
	//}
	//// 创建一个日志对象
	//debugLog := log.New(logFile,"[Debug]",log.LstdFlags)
	//debugLog.Println("A debug message here")
	////配置一个日志格式的前缀
	//debugLog.SetPrefix("[Info]")
	//debugLog.Println("A Info Message here ")
	////配置log的Flag参数
	//debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	//debugLog.Println("A different prefix")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	//4.当访问 localhost/healthz 时，应返回200
	io.WriteString(w, "200")
	w.WriteHeader(200)
}

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/task1/", task1)
	http.HandleFunc("/task2/", task2)
	http.HandleFunc("/task3/", task3)
	http.HandleFunc("/healthz/", healthz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
