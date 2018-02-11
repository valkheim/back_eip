package main

func main() {
  a := API{Address:"api:443", Timeout:15, Router:nil}
  a.Run()
}
