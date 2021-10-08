package main

import "net/http"

type HandlerBaseOnMap struct {
	handlers map[string]func(ctx *Context)
}

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
