package server

import (
	"net/http"

	"ck/dispatcher"
	"ck/handlers"
)

func New(d *dispatcher.Dispatcher) *http.Server {
	return &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: handlers.New(d),
	}
}
