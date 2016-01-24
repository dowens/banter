# Changes to banter

## v0.1.1
* Fixing casting issue for constructor middleware and adding case to
  router_test.go.

## v0.1.0
* Rewrote middleware interface to function like and use
  [Alice](https://github.com/justinas/alice).
* Changed the handler signature to match http.HandlerFunc for consistency.
* Moved URL params out of (removed) context struct and into URL query.

## v0.0.1
* Initial version
