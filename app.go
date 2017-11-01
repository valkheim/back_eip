package main

import (
    "log"
    "time"
    "net/http"

    "github.com/gorilla/mux"
)

type App struct {
  Address string
  Timeout time.Duration
  Router *mux.Router
}

func (app *App) Run() {
  app.Router = mux.NewRouter()
  srv := &http.Server {
    Handler:      app.Router,
    Addr:         app.Address,
    WriteTimeout: app.Timeout * time.Second,
    ReadTimeout:  app.Timeout * time.Second,
  }

  app.initializeRoutes()

  log.Fatal(srv.ListenAndServe())
}
