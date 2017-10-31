package main

import (
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
}

/* Get client informations */
func getClientHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintf(w, "Vars %v\n", vars)
}

/* Create client */
func createClientHandler(f func(client Client)) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    var client Client
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
      panic(err)
    }
    if err := r.Body.Close(); err != nil {
      panic(err)
    }
    if err := json.Unmarshal(body, &client); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusUnprocessableEntity)
      if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
      }
    }

    f(client)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(client); err != nil {
      panic(err)
    }
  }
}

func createClient(client Client) {
  fmt.Printf("New client %v\n", client)
}
