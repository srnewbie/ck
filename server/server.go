package server

import (
	"ck/dispatcher"
	"ck/handlers"
	"net/http"
)

func New(d *dispatcher.Dispatcher) *http.Server {
	return &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: handlers.New(d),
	}
}
