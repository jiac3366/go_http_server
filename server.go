package main

import (
	"net/http"
)

//这个文件对GO原生的 注册路由和监听端口逻辑绑定在一起

// Server 下面这个Server可以不放这里 可以放其他文件 --这个不引入包 但实现了接口 挺屌
type Server interface {
	Route(pattern string, handleFunc func(ctx *Context)) // http.HandlerFunc --> func (ctx *Context)是为了控制context的创建，所以你传入的方法接受的参数是Context
	Start(adress string) error
}

// sdkHttpServer 基于http库实现
type sdkHttpServer struct {
	Name string
}

// Route 路由注册
func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *Context)) { //http.HandlerFunc --> func (ctx *Context)
	// 唯一细节相关的就是http.HandleFunc
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		// handleFunc --> func(writer http.ResponseWriter, request *http.Request) {....}
		//闭包
		// 不想让Server知道ctx内部创建细节
		ctx := NewContext(writer, request)
		handleFunc(ctx)
	})
}

// Start 监听端口
func (s *sdkHttpServer) Start(adress string) error {
	return http.ListenAndServe(adress, nil)
}

// NewServer 创建完struct会对应创建一个创建的方法（相当于构造函数）
func NewServer() Server {
	return &sdkHttpServer{}
}

// NewHttpServer 让用户第一次进去不知道我sdkHttpServer的源码
func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
