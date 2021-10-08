package main

import (
	"net/http"
)

//这个文件对GO原生的 注册路由和监听端口逻辑绑定在一起

// Server 下面这个Server可以不放这里 可以放其他文件 --这个不引入包 但实现了接口 挺屌
type Server interface {
	Route(method string, pattern string, handleFunc func(ctx *Context)) // http.HandlerFunc --> func (ctx *Context)是为了控制context的创建，所以你传入的方法接受的参数是Context
	Start(adress string) error
}

// sdkHttpServer 基于http库实现
type sdkHttpServer struct {
	Name    string
	handler *HandlerBaseOnMap //
}

// Route 路由注册
func (s *sdkHttpServer) Route(
	method string, // 添加Restful API
	pattern string,
	handleFunc func(ctx *Context)) { //http.HandlerFunc --> func (ctx *Context)

	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handleFunc
	//// 唯一细节相关的就是http.HandleFunc
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	// handleFunc --> func(writer http.ResponseWriter, request *http.Request) {....}
	//	//闭包
	//	// 不想让Server知道ctx内部创建细节
	//	ctx := NewContext(writer, request)
	//	handleFunc(ctx)
	//})
}

// Start 监听端口
func (s *sdkHttpServer) Start(adress string) error {
	// 这里很重要：移到后面统一给所有的路由创建系统上下文（writer and request）
	// 可以注意一下NewContext的位置从Route()到ServeHTTP()
	http.Handle("/", s.handler)
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
