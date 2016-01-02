# banter
**httprouter + interface for middleware = go web framework**


## Usage

Configure middleware and route handlers using the same function signature:
```go
func middleware1(res http.ResponseWriter, req *http.Request, context banter.Context) {
  // Do something...
  // Call the next handler.
  banter.Next(res, req, context)
}

func handler(res http.ResponseWriter, req *http.Request, context banter.Context) {
  // Print the ID parameter to the response.
  fmt.Fprintf(res, "The resource ID is: %s", context.Params.ByName("id"))
}
```

Create a router, configure it, and run it:
```go
// Create a router.
router := banter.Router()

// Configure global middleware.
router.Use(middlware1)

// Configure a route with route-specific middleware and final handler.
router.GET("/thing/:id", middleware2, middleware3, handler)

// Run the server.
http.ListenAndServe(":8080", router)
```
