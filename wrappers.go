package main

import (
  "net/http"
  "fmt"
  "io"
  "io/ioutil"
  "reflect"
  "encoding/json"
)

/* Create new item */
func createHandler(f func(T interface{}) error, T interface{}) func(w http.ResponseWriter, r *http.Request) {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    newItem := reflect.ValueOf(T).Interface() // get underlying type as run-time under a value and transform it as an interface
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
      panic(err)
    }
    if err := r.Body.Close(); err != nil {
      panic(err)
    }
    if err := json.Unmarshal(body, &newItem); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusUnprocessableEntity)
      if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
      }
    }
    if err := f(newItem) ; err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusConflict)
      fmt.Println(err)
    } else {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusCreated)
      if err := json.NewEncoder(w).Encode(newItem); err != nil {
        panic(err)
      }
    }
  })
}

func checkHeaders(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    contentType := r.Header.Get("Content-Type")
    if contentType == "" || contentType != "application/json" {
      w.WriteHeader(http.StatusUnsupportedMediaType)
    } else {
      inner.ServeHTTP(w, r)
    }
  })
}

func addDefaultHeaders(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
    inner.ServeHTTP(w, r)
  })
}
