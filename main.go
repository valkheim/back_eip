package main

func main() {
  a := Api{Address:"api:443", Timeout:15, Router:nil}
  a.Run()
}
