package http_order

import "github.com/go-chi/chi/v5"

type Handler interface {
	Register(chi.Router)
}
