package main

import (
  "reflect"
  "fmt"
)

type Partner struct {
  Id    int     `json:"id"`
  Name  string  `json:"name"`
}

/* Create partner (register in db/store) */
func createPartner(T interface{}) error {
  partner := *(reflect.ValueOf(T).Elem()).Addr().Interface().(*Partner)
  fmt.Printf("New partner type %T\n", partner)
  fmt.Printf("New partner %v\n", partner)
  return nil
}
