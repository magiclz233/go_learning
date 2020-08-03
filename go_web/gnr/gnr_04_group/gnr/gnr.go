package gnr

import (
	"log"
	"net/http"
)

// HandlerFunc 定义gnr使用的请求处理程序 类似 http.HandlerFunc
type HandlerFunc func(*Context)

// 实现ServeHTTP的接口
type (
	RouterGroup struct {
		prefix      string
		middleWares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
)

// 构造函数  初始化gnr
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (r *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := r.prefix + comp
	log.Printf("Route %4s - %s\n", method, pattern)
	r.engine.router.addRoute(method, pattern, handler)
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// GET定义添加GET请求的方法
func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

// POST定义添加POST请求的方法
func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}

// 启动http server
func (e *Engine) Run(add string) error {
	return http.ListenAndServe(add, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
