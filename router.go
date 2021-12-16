package gss

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, pattern string, handle HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handle
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c.Writer, c.Req)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%s \n", c.Path)
	}
}
