package main

func (app *App) initializeRoutes() {
  app.Router.HandleFunc("/client/{id:[0-9]+}", getClientHandler).Methods("GET")
  app.Router.HandleFunc("/client", createHandler(createClient, &Client{})).Methods("POST")
}
