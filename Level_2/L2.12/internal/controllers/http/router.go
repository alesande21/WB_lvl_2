package http

import (
	"net/http"
	"regexp"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(method, path string, handler http.Handler) {
	r.routes = append(r.routes, Route{
		Method:  method,
		Pattern: path,
		Handler: handler,
	})
}

func (r *Router) getHandler(method, path string) http.Handler {
	for _, route := range r.routes {
		re := regexp.MustCompile(route.Pattern)
		if route.Method == method && re.MatchString(path) {
			return route.Handler
		}
	}
	return http.NotFoundHandler()
}
