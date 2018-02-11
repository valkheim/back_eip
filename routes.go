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
	{"Get client infos", "GET", "/client/{id:[0-9]+}", getClientHandler},
	{"Delete client", "DELETE", "/client/{id:[0-9]+}", deleteClientHandler},
	{"Create client", "POST", "/client", createHandler(createClient, &Client{})},
	{"Create partner", "POST", "/partner", createHandler(createPartner, &Partner{})},
}
