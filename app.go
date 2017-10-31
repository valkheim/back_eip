package main

import (
    "log"
    "time"
    "net/http"
    "encoding/json"

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

func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
