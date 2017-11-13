package main

import (
  "net/http"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

var routes = []Route{
  Route{"Get client infos", "GET", "/client/{id:[0-9]+}", getClientHandler, },
  Route{"Delete client", "DELETE", "/client/{id:[0-9]+}", deleteClientHandler, },
  Route{"Create client", "POST", "/client", createHandler(createClient, &Client{}), },
  Route{"Create partner", "POST", "/partner", createHandler(createPartner, &Partner{}), },
}
