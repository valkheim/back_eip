package main

import (
  "reflect"
  "net/http"
  "fmt"
  "errors"

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

/* Delete client (remove from db/store) */
func deleteClientHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Printf("Deleting client using vars %v\n", vars)
  w.WriteHeader(http.StatusNoContent)
}

/* Create client (register in db/store) */
func createClient(T interface{}) error {
  // valueof : get underlying value of interface T
  // elem : return the value that the interface contains
  // *().Addr Obtain a pointer to the value
  // cast/type assertion ?
  //client := *(reflect.ValueOf(T).Elem()).Addr().Interface().(*Client)
  client := (reflect.ValueOf(T)).Interface().(*Client)
  fmt.Printf("New client type %T\n", client)
  fmt.Printf("New client %v\n", client)
  if 1 == 2 {
    return errors.New("Cannot create client")
  }
  return nil
}
