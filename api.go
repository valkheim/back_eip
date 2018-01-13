package main

import (
  "log"
  "time"
  "net/http"
  "crypto/tls"

  "github.com/gorilla/mux"
)

type Api struct {
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
  http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func (app *Api) Run() {
  go http.ListenAndServe(":80", http.HandlerFunc(redirect))
  if app.Router == nil {
    app.Router = NewRouter()
  }

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

  log.Fatal(srv.ListenAndServeTLS("auth/server.crt", "auth/server.key"))
}
