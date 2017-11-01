package main

import (
  "log"
  "io"
  "io/ioutil"
  "time"
  "net/http"
  "reflect"
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

/* Create new item */
func createHandler(f func(T interface{}), T interface{}) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    newItem := reflect.ValueOf(T).Interface()

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
      panic(err)
    }
    if err := r.Body.Close(); err != nil {
      panic(err)
    }
    /* We got to write our own Unmarshal to check for required fields */
    if err := json.Unmarshal(body, &newItem); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusUnprocessableEntity)
      if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
      }
    }

    f(newItem)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(newItem); err != nil {
      panic(err)
    }
  }
}
