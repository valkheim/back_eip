package main

//import (
//  "fmt"
//  "./store"
//)

func main() {
  //fmt.Print(store.GetAllKeys(), "\n")
  a := Api{Address:"127.0.0.1:443", Timeout:15, Router:nil}
  a.Run()
}
