package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

// NewRouter creates a gorilla/mux router and binds routes and its middlewares.
func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler
    handler = route.HandlerFunc
    handler = addDefaultHeaders(handler, route.Name)
    handler = checkHeaders(handler, route.Name)
    handler = Logger(handler, route.Name)
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }
  return router
}
