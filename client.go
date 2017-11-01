package main

import (
  "reflect"
  "net/http"
  "fmt"
  //"errors"

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

/* Create client (register in db/store) */
func createClient(T interface{}) error {
  client := *(reflect.ValueOf(T).Elem()).Addr().Interface().(*Client)
  fmt.Printf("New client type %T\n", client)
  fmt.Printf("New client %v\n", client)
  //return errors.New("np")
  return nil
}
