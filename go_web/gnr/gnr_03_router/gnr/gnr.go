package gnr

import (
	"log"
	"net/http"
)

// HandlerFunc 定义gnr使用的请求处理程序 类似 http.HandlerFunc
type HandlerFunc func(*Context)

// 实现ServeHTTP的接口
type Engine struct {
	router *router
}

// 构造函数  初始化gnr
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s\n", method, pattern)
	e.router.addRoute(method, pattern, handler)
}

// GET定义添加GET请求的方法
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST定义添加POST请求的方法
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// 启动http server
func (e *Engine) Run(add string) error {
	return http.ListenAndServe(add, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
