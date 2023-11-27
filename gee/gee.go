package gee

import (
	"net/http"
)

type RouterGroup struct {
	middlewares []HandleFunc
	prefix      string
	parent      *RouterGroup
	engine      *Engine
	middleware  []*HandleFunc
}

type HandleFunc func(c *Context)
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
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

func (group *RouterGroup) addRoute(method string, pattern string, handle HandleFunc) {
	//group

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
	c := NewContext(w, req)
	e.router.handle(c)
}
