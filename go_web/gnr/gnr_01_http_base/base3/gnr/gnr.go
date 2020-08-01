package gnr

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc 定义gnr使用的请求处理程序 类似 http.HandlerFunc
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

// 实现ServeHTTP的接口
type Engine struct {
	router map[string]HandlerFunc
}

// 构造函数  初始化gnr
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s\n", method, pattern)
	e.router[key] = handler
}

// GET定义添加GET请求的方法
func (e Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

// POST定义添加POST请求的方法
func (e Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

// 启动http server
func (e Engine) Run(add string) error {
	return http.ListenAndServe(add, e)
}

func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUND %s\n", req.URL)
	}
}
