package main

import (
	"fmt"
	"net/http"
	"reflect"

	"./store"
	"github.com/gorilla/mux"
)

// Client represents a customer of partner stores
type Client struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Get client informations
func getClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Vars %v\n", vars)
}

// Delete client (remove from db/store)
func deleteClientHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("Deleting client using vars %v\n", vars)
	w.WriteHeader(http.StatusNoContent)
}

// Create client (register in db/store)
func createClient(T interface{}) error {
	c := (reflect.ValueOf(T)).Interface().(*Client)
	req := fmt.Sprintf("HMSET client:%d id %d name %s age %d", c.ID, c.ID, c.Name, c.Age)
	return store.Exec(req)
}
