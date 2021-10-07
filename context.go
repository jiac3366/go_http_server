package main

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//这个文件把response和request逻辑绑定 使得server struct可以调用 ReadJson 和 WriteJson

// 目前来看 Context不需要扩充什么功能 所以用结构体

type Context struct {
	W http.ResponseWriter // ResponseWriter没有*是因为 返回的是一个接口
	R *http.Request       // Request返回的是结构体 不怎么再改变了
}

func (c *Context) ReadJson(req interface{}) error {
	// ReadJson要读出body interface{}相当于object
	// go一般传一个空的结构体进来 然后修改里面的数据并且返回是否成功
	r := c.R
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}

	return nil

}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson)
	return err

}

func (c *Context) BadReqJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

// 不希望Server理解如何创建context的细节

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}
}
