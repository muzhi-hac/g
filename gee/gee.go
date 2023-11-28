package gee

import (
	"log"
	"net/http"
	"strings"
)

type RouterGroup struct {
	//middlewares []HandleFunc
	prefix      string
	parent      *RouterGroup
	engine      *Engine
	middlewares []HandleFunc
}

type HandleFunc func(c *Context)
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func (group *RouterGroup) Use(midd ...HandleFunc) {
	group.middlewares = append(group.middlewares, midd...)

}
func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	//return &Engine{router: NewRouter()}
	return engine
}
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{prefix: group.prefix + prefix,

		engine: group.engine}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handle HandleFunc) {
	//group
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	//group.engine.router.addRoute(method, pattern, handler)
	group.engine.router.addRouter(method, pattern, handle)

}

//	func (e *Engine) addRouter(method string, pattern string, handle HandleFunc) {
//		router := method + "-" + pattern
//		e.router[router] = handle
//		//return e
//	}
//
// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}
func (e *Engine) Run(address string) error {
	return http.ListenAndServe(address, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandleFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}

	}
	c := NewContext(w, req)
	c.handlers = middlewares
	c.engine = e
	//engine.router.handle(c)
	e.router.handle(c)
}
