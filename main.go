package main

import (
  "fmt"
  "strconv"
)

func main() {
  var c Config
  if e := c.Get("backend.toml"); e != nil {
    fmt.Println(e)
  }
  fmt.Println(c.General.Version)
  fmt.Println(c.General.Address)
  fmt.Println(c.General.Port)
  fmt.Printf("dt %T\n" ,c.General.Timeout.Duration)
  a := App{Address:c.General.Address+":"+strconv.Itoa(c.General.Port), Timeout:15}
  a.Run()
}
