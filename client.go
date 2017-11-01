package main

import (
  "reflect"
  "net/http"
  "fmt"
  "io"
  "io/ioutil"
  "encoding/json"

  "github.com/gorilla/mux"
)

type Client struct {
  Id    int     `json:"id"`
  Name  string  `json:"name"`
  Age   int     `json:"age"`
}

/* Get client informations */
func getClientHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintf(w, "Vars %v\n", vars)
}

/* Create new item */
func createHandler(f func(T interface{}), T interface{}) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    newItem := reflect.ValueOf(T).Interface()

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

    f(newItem)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(newItem); err != nil {
      panic(err)
    }
  }
}

func createClient(T interface{}) {
  v := reflect.ValueOf(T).Elem()
  ptr := v.Addr().Interface().(*Client)
  client := *ptr
  fmt.Printf("New client type %T\n", client)
  fmt.Printf("New client %v\n", client)
}
