package main

import "net/http"

//这个文件 是为了

//要往里面加方法了, 第三节课：把Handler也抽象
type Handler interface {
	http.Handler //组合？
	Route(method string, pattern string,
		handleFunc func(ctx *Context))
}

type HandlerBaseOnMap struct {
	handlers map[string]func(ctx *Context)
}

// 把server的Route() 搬了过来
func (h *HandlerBaseOnMap) Route(
	method string, // 添加Restful API
	pattern string,
	handleFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handleFunc

}

// ServeHTTP 为已经注册的路由创建上下文
func (h *HandlerBaseOnMap) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	// 如果ok 说明已经注册过了
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
}

func (h *HandlerBaseOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerBaseOnMap() Handler {
	return &HandlerBaseOnMap{
		handlers: make(map[string]func(c *Context)),
	}
}
