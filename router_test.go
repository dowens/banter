package banter

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
	"testing"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func TestRouter(t *testing.T) {
	hasMiddlewareBeenCalled := false
	hasHandlerBeenCalled := false

	// Configure the router.
	router := Router()
	router.GET(
		"/user/:name",
		func(res http.ResponseWriter, req *http.Request, context Context) {
			hasMiddlewareBeenCalled = true
			Next(res, req, context)
		},
		func(res http.ResponseWriter, req *http.Request, context Context) {

			// Verify middleware has been called before this Handler executes.
			if !hasMiddlewareBeenCalled {
				t.Fatal("Middleware not called.")
			}

			hasHandlerBeenCalled = true

			// Verify Params are correct.
			want := httprouter.Params{httprouter.Param{"name", "gopher"}}
			if !reflect.DeepEqual(context.Params, want) {
				t.Fatalf("Wrong wildcard values: want %v, got %v", want, context.Params)
			}

		},
	)

	// Make the test request.
	req, _ := http.NewRequest("GET", "/user/gopher", nil)
	router.ServeHTTP(new(mockResponseWriter), req)

	// Verify hanlder has been called.
	if !hasHandlerBeenCalled {
		t.Fatal("Handler not called.")
	}
}

func TestMiddleware(t *testing.T) {
	hasMiddlewareBeenCalled := false
	hasHandlerBeenCalled := false

	// Create the router.
	router := Router()

	// Configure the middleware.
	router.Use(func(res http.ResponseWriter, req *http.Request, context Context) {
		hasMiddlewareBeenCalled = true
		Next(res, req, context)
	})

	// Configure the test route.
	router.GET(
		"/user/:name",
		func(res http.ResponseWriter, req *http.Request, context Context) {

			// Verify middleware has been called before this Handler executes.
			if !hasMiddlewareBeenCalled {
				t.Fatal("Middleware not called.")
			}

			hasHandlerBeenCalled = true

		},
	)

	// Make the test request.
	req, _ := http.NewRequest("GET", "/user/gopher", nil)
	router.ServeHTTP(new(mockResponseWriter), req)

	// Verify hanlder has been called.
	if !hasHandlerBeenCalled {
		t.Fatal("Handler not called.")
	}
}
