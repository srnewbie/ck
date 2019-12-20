package handlers

import (
	"ck/dispatcher"
	"ck/models"
	"encoding/json"
	"net/http"
)

type Handlers struct {
	Dispatcher *dispatcher.Dispatcher
}

func New(d *dispatcher.Dispatcher) *Handlers {
	return &Handlers{
		Dispatcher: d,
	}
}

func (h *Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("OK Get"))
	case "POST":
		order := &models.Order{}
		err := json.NewDecoder(r.Body).Decode(order)
		if err != nil {
			w.Write([]byte("Bad Request"))
			panic(err)
		}
		h.Dispatcher.Push(order)
		w.Write([]byte("OK Post"))
	default:
		w.Write([]byte("Wrong API type"))
	}
}
