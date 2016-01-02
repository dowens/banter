package banter

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
)

type Handler func(http.ResponseWriter, *http.Request, Context)

type Context struct {
	Params httprouter.Params
	Handlers []Handler
}

type HttpRouter struct {
  Router *httprouter.Router
}

/*
Router TODO
*/
func Router() *HttpRouter {
  return &HttpRouter{Router: httprouter.New()}
}

// func (r *httprouter.Router) Use() {
//
// }

// GET is a shortcut for router.AddHandler("GET", path, handle)
func (r *HttpRouter) GET(path string, handlers ...Handler) {
	r.AddHandlers("GET", path, handlers)
}

// HEAD is a shortcut for router.AddHandler("HEAD", path, handle)
func (r *HttpRouter) HEAD(path string, handlers ...Handler) {
	r.AddHandlers("HEAD", path, handlers)
}

// OPTIONS is a shortcut for router.AddHandler("OPTIONS", path, handle)
func (r *HttpRouter) OPTIONS(path string, handlers ...Handler) {
	r.AddHandlers("OPTIONS", path, handlers)
}

// POST is a shortcut for router.AddHandler("POST", path, handle)
func (r *HttpRouter) POST(path string, handlers ...Handler) {
	r.AddHandlers("POST", path, handlers)
}

// PUT is a shortcut for router.AddHandler("PUT", path, handle)
func (r *HttpRouter) PUT(path string, handlers ...Handler) {
	r.AddHandlers("PUT", path, handlers)
}

// PATCH is a shortcut for router.AddHandler("PATCH", path, handle)
func (r *HttpRouter) PATCH(path string, handlers ...Handler) {
	r.AddHandlers("PATCH", path, handlers)
}

// DELETE is a shortcut for router.AddHandler("DELETE", path, handle)
func (r *HttpRouter) DELETE(path string, handlers ...Handler) {
	r.AddHandlers("DELETE", path, handlers)
}

/*
AddHandlers TODO
*/
func (r *HttpRouter) AddHandlers(
  method string,
  path string,
  handlers []Handler,
) {
  r.Router.Handle(
    method,
    path,
    func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
      Next(res, req, Context{Params: params, Handlers: handlers})
    },
  )
}

/*
Next TODO
*/
func Next(res http.ResponseWriter, req *http.Request, context Context) {
	if len(context.Handlers) > 0 {
		handler := context.Handlers[0]
		context.Handlers = context.Handlers[1:]
		handler(res, req, context)
	}
}

/*
ServeHTTP TODO
*/
func (r *HttpRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  r.Router.ServeHTTP(res, req)
}
