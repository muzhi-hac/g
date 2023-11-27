package gee

import (
	"net/http"
	"strings"
)

type router struct {
	handler map[string]HandleFunc
	root    map[string]*node
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, pattern string, handle HandleFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.root[method]; !ok {
		r.root[method] = &node{}
	}
	r.root[method].Insert(pattern, parts, 0)
	r.handler[key] = handle

}
func (r *router) getRouter(method string, pattern string) (*node, map[string]string) {
	n := r.root[method].Search(parsePattern(pattern), 0)
	params := make(map[string]string)
	if n != nil {
		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = parsePattern(pattern)[i]
			} else if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(parsePattern(pattern)[i:], "/")
				break
			}
			return n, params
		}
	}
	return nil, nil

}

//func (group *RouterGroup) Group(prefix string) {
//	engine := group.engine
//	newGroup := &RouterGroup{prefix: group.prefix + prefix,
//
//		engine: group.engine}
//
//	engine.group = append(engine.routerGroup, newGroup)
//
//}

func NewRouter() *router {
	return &router{handler: make(map[string]HandleFunc),
		root: make(map[string]*node)}
}

//	func (r *router) addRouter(method string, pattern string, handler HandleFunc) {
//		path := method + "-" + pattern
//		r.handler[path] = handler
//
// }
func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + c.Path
		c.Params = params
		c.handlers = append(c.handlers, r.handler[key])
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND")

	}

}
