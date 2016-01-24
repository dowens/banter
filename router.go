package banter

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	// "log"
	"net/http"
)

/*
HTTPHandler TODO
*/
type HTTPHandler struct {
	Handler http.Handler
}

/*
HTTPRouter TODO
*/
type HTTPRouter struct {
	Router *httprouter.Router
	Chain  alice.Chain
}

/*
Constructor TODO
*/
func (h HTTPHandler) Constructor(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Handler.ServeHTTP(w, r)
		handler.ServeHTTP(w, r)
	})
}

/*
Router TODO
*/
func Router() *HTTPRouter {
	return &HTTPRouter{Router: httprouter.New(), Chain: alice.New()}
}

/*
Chain TODO
*/
func Chain(chain *alice.Chain, middleware []interface{}) alice.Chain {
	var c alice.Chain
	var chainedMiddleware []alice.Constructor
	for _, m := range middleware {
		if constructor, isCon := m.(func(http.Handler) http.Handler); isCon {
			chainedMiddleware = append(chainedMiddleware, constructor)
		} else if h, isH := m.(func(http.ResponseWriter, *http.Request)); isH {
			mwHandler := HTTPHandler{http.HandlerFunc(h)}
			chainedMiddleware = append(chainedMiddleware, mwHandler.Constructor)
		} else {
			panic("Middleware must match either the alice.Constructor or the " +
				"http.HandlerFunc type")
		}
	}
	c = alice.New(chainedMiddleware...)
	if chain != nil {
		c.Extend((*chain))
	}
	return c
}

/*
Use TODO
*/
func (r *HTTPRouter) Use(middleware ...interface{}) {
	r.Chain = Chain(&r.Chain, middleware)
}

// GET is a shortcut for router.AddHandler("GET", path, handle)
func (r *HTTPRouter) GET(path string, middleware ...interface{}) {
	r.AddHandlers("GET", path, middleware)
}

// HEAD is a shortcut for router.AddHandler("HEAD", path, handle)
func (r *HTTPRouter) HEAD(path string, middleware ...interface{}) {
	r.AddHandlers("HEAD", path, middleware)
}

// OPTIONS is a shortcut for router.AddHandler("OPTIONS", path, handle)
func (r *HTTPRouter) OPTIONS(path string, middleware ...interface{}) {
	r.AddHandlers("OPTIONS", path, middleware)
}

// POST is a shortcut for router.AddHandler("POST", path, handle)
func (r *HTTPRouter) POST(path string, middleware ...interface{}) {
	r.AddHandlers("POST", path, middleware)
}

// PUT is a shortcut for router.AddHandler("PUT", path, handle)
func (r *HTTPRouter) PUT(path string, middleware ...interface{}) {
	r.AddHandlers("PUT", path, middleware)
}

// PATCH is a shortcut for router.AddHandler("PATCH", path, handle)
func (r *HTTPRouter) PATCH(path string, middleware ...interface{}) {
	r.AddHandlers("PATCH", path, middleware)
}

// DELETE is a shortcut for router.AddHandler("DELETE", path, handle)
func (r *HTTPRouter) DELETE(path string, middleware ...interface{}) {
	r.AddHandlers("DELETE", path, middleware)
}

/*
AddHandlers TODO
*/
func (r *HTTPRouter) AddHandlers(
	method string,
	path string,
	middleware []interface{},
) {
	r.Router.Handle(
		method,
		path,
		func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
			// Add URL params as query params so we can stay consistent with the
			// interface.
			for _, param := range params {
				values := req.URL.Query()
				values.Add(param.Key, param.Value)
				req.URL.RawQuery = values.Encode()
			}

			// Create chain of middleware.
			c := Chain(nil, middleware)

			// Serve chain.
			c.Then(nil).ServeHTTP(res, req)
		},
	)
}

/*
ServeHTTP TODO
*/
func (r *HTTPRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	r.Chain.Then(r.Router).ServeHTTP(res, req)
}
