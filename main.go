package main

func main() {
  a := Api{Address:"api:443", Timeout:15, Router:nil}
  //a := Api{Address:"127.0.0.1:443", Timeout:15, Router:nil}
  a.Run()
}
