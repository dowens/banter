# banter
**[httprouter](https://github.com/julienschmidt/httprouter) +
[alice](https://github.com/justinas/alice)-like middleware = go web framework**


## Usage

Your middleware and route handlers can have the same function signature:
```go
func customMiddleware(res http.ResponseWriter, req *http.Request) {
  // Do something before the final handler executes...
}

func handler(res http.ResponseWriter, req *http.Request) {
  // URL params (like :id) are added as query params
  fmt.Fprintf(res, "The resource ID is: %s", req.Query().Get("id"))
}
```

Create a router, configure it, and run it:
```go
// Create a router.
router := banter.Router()

// Configure global middleware.
router.Use(cors.Default().Handler, logger.New().Handler)

// Configure a route with route-specific middleware and handlers.
router.GET("/thing/:id", customMiddleware, nosurf.NewPure, handler)

// Run the server.
http.ListenAndServe(":8080", router)
```
