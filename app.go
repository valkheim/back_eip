package main

import (
  "fmt"
  "log"
  "io"
  "io/ioutil"
  "time"
  "net/http"
  "crypto/tls"
  "reflect"
  "encoding/json"

  "github.com/gorilla/mux"
)

type App struct {
  Address string
  Timeout time.Duration
  Router *mux.Router
}


func redirect(w http.ResponseWriter, req *http.Request) {
  target := "https://" + req.Host + req.URL.Path
  if len(req.URL.RawQuery) > 0 {
    target += "?" + req.URL.RawQuery
  }
  log.Printf("Redirect to: %s", target)
  http.Redirect(w, req, target,
  http.StatusTemporaryRedirect)
}

func (app *App) Run() {
  go http.ListenAndServe(":80", http.HandlerFunc(redirect))
  app.Router = mux.NewRouter()

  cfg := &tls.Config{
    MinVersion:               tls.VersionTLS12,
    CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
    PreferServerCipherSuites: true,
    CipherSuites: []uint16{
      tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
      tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
      tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
      tls.TLS_RSA_WITH_AES_256_CBC_SHA,
    },
  }

  srv := &http.Server {
    Handler:      app.Router,
    Addr:         app.Address,
    WriteTimeout: app.Timeout * time.Second,
    ReadTimeout:  app.Timeout * time.Second,
    MaxHeaderBytes: 1 << 20, // 1 MB (default value)
    TLSConfig:    cfg,
    TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
  }

  app.initializeRoutes()

  log.Fatal(srv.ListenAndServeTLS("auth/server.crt", "auth/server.key"))
}

/* Create new item */
func createHandler(f func(T interface{}) error, T interface{}) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    newItem := reflect.ValueOf(T).Interface()

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
      panic(err)
    }
    if err := r.Body.Close(); err != nil {
      panic(err)
    }
    if err := json.Unmarshal(body, &newItem); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusUnprocessableEntity)
      if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
      }
    }
    if err := f(newItem) ; err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusConflict)
      fmt.Println(err)
    } else {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusCreated)
      if err := json.NewEncoder(w).Encode(newItem); err != nil {
        panic(err)
      }
    }
  }
}
