package router

import (
	"fmt"
	"net/http"

	"github.com/faoztas/gocache/utils/errors"
)

type Router interface {
	GET(prefix string, handler http.HandlerFunc)
	PUT(prefix string, handler http.HandlerFunc)
	PATCH(prefix string, handler http.HandlerFunc)
	POST(prefix string, handler http.HandlerFunc)
	DELETE(prefix string, handler http.HandlerFunc)
	Check(request *http.Request) (error, int)
	GetMethods() *map[string]map[string]http.HandlerFunc
}

type router struct {
	Routers map[string]map[string]http.HandlerFunc
}

func NewRouter() Router {
	return &router{
		Routers: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *router) GET(prefix string, handler http.HandlerFunc) {
	r.setMethod(http.MethodGet)
	r.Routers[http.MethodGet][prefix] = handler
}

func (r *router) PUT(prefix string, handler http.HandlerFunc) {
	r.setMethod(http.MethodPut)
	r.Routers[http.MethodPut][prefix] = handler
}

func (r *router) PATCH(prefix string, handler http.HandlerFunc) {
	r.setMethod(http.MethodPatch)
	r.Routers[http.MethodPatch][prefix] = handler
}

func (r *router) POST(prefix string, handler http.HandlerFunc) {
	r.setMethod(http.MethodPost)
	r.Routers[http.MethodPost][prefix] = handler
}

func (r *router) DELETE(prefix string, handler http.HandlerFunc) {
	r.setMethod(http.MethodDelete)
	r.Routers[http.MethodDelete][prefix] = handler
}

func (r *router) Check(request *http.Request) (error, int) {
	prefix := request.URL.Path

	match := false
	for method, list := range r.Routers {
		if method == request.Method {
			if list[prefix] == nil {
				return fmt.Errorf(errors.PrefixNotFound, request.RequestURI, request.Method), http.StatusNotFound
			}
			match = true
		}
	}
	if !match {
		return fmt.Errorf(errors.MethodNotAllowed), http.StatusMethodNotAllowed
	}
	return nil, http.StatusOK
}

func (r *router) GetMethods() *map[string]map[string]http.HandlerFunc {
	return &r.Routers
}

func (r *router) setMethod(method string) {
	if r.Routers[method] == nil {
		r.Routers[method] = make(map[string]http.HandlerFunc)
	}
}
