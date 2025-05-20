package http

import (
	"hexademo/internal/domain"
	"net/http"
)

// Handler bridges HTTP requests to the Greeter domain.
type Handler struct {
	greeter domain.Greeter
}

func NewHandler(g domain.Greeter) *Handler {
	return &Handler{greeter: g}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	greeting := h.greeter.Greet(name)
	_, _ = w.Write([]byte(greeting))
}
