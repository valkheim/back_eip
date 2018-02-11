package main

import (
	"./store"
	"fmt"
	"reflect"
)

// Partner represents a GroomShop partner store.
type Partner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

/* Create partner (register in db/store) */
func createPartner(T interface{}) error {
	//partner := *(reflect.ValueOf(T).Elem()).Addr().Interface().(*Partner)
	p := (reflect.ValueOf(T)).Interface().(*Partner)
	req := fmt.Sprintf("HMSET partner:%d id %d name %s", p.ID, p.ID, p.Name)
	return store.Exec(req)
}
