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

  router := Router()
  router.GET(
    "/user/:name",
    func(res http.ResponseWriter, req *http.Request, context Context) {
      hasMiddlewareBeenCalled = true
      Next(res, req, context)
    },
    func(res http.ResponseWriter, req *http.Request, context Context) {
      hasHandlerBeenCalled = true

      want := httprouter.Params{httprouter.Param{"name", "gopher"}}
      if !reflect.DeepEqual(context.Params, want) {
        t.Fatalf("Wrong wildcard values: want %v, got %v", want, context.Params)
      }
    },
  )

  req, _ := http.NewRequest("GET", "/user/gopher", nil)
  router.ServeHTTP(new(mockResponseWriter), req)

  if !hasMiddlewareBeenCalled {
    t.Fatal("Middleware not called.")
  }
  if !hasHandlerBeenCalled {
    t.Fatal("Handler not called.")
  }
}
