package main

import "net/http"

func addDefaultHeaders(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    /* Gotta check preloading eligibility */
    w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
    f(w, r)
  }
}

func (app *App) initializeRoutes() {
  app.Router.HandleFunc("/client/{id:[0-9]+}", addDefaultHeaders(getClientHandler)).Methods("GET")
  app.Router.HandleFunc("/client/{id:[0-9]+}", addDefaultHeaders(deleteClientHandler)).Methods("DELETE")
  app.Router.HandleFunc("/client", addDefaultHeaders(createHandler(createClient, &Client{}))).Methods("POST")
  app.Router.HandleFunc("/partner", addDefaultHeaders(createHandler(createPartner, &Partner{}))).Methods("POST")

  app.Router.HandleFunc("/debugssl", func(w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
    w.Write([]byte("U got a TLS line genius !!1\n"))
  })

}
