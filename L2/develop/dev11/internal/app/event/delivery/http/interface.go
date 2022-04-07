package event

import "net/http"

type Handler interface {
	Register(*http.ServeMux)
}
