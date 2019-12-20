package handlers

import (
	"encoding/json"
	"net/http"

	"ck/dispatcher"
	"ck/models"
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
		w.Write([]byte("Not implemented"))
	case "POST":
		orders := []*models.Order{}
		err := json.NewDecoder(r.Body).Decode(&orders)
		if err != nil {
			w.Write([]byte("Bad Request"))
			panic(err)
		}
		for _, o := range orders {
			h.Dispatcher.Push(o)
		}
		w.Write([]byte("OK Post"))
	default:
		w.Write([]byte("Wrong API type"))
	}
}
