package main

import (
  "reflect"
  "net/http"
  "fmt"

  "github.com/gorilla/mux"
  "./store"
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

/* Delete client (remove from db/store) */
func deleteClientHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Printf("Deleting client using vars %v\n", vars)
  w.WriteHeader(http.StatusNoContent)
}

/* Create client (register in db/store) */
func createClient(T interface{}) error {
  c := (reflect.ValueOf(T)).Interface().(*Client)
  req := fmt.Sprintf("HMSET client:%d id %d name %s age %d", c.Id, c.Id, c.Name, c.Age)
  return store.Exec(req)
}
