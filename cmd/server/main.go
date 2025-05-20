package main

import (
	"log"
	"net/http"

	httpadapter "hexademo/internal/adapter/http"
	"hexademo/internal/service"
)

// listenAndServe is defined as a variable so it can be replaced in tests.
var listenAndServe = http.ListenAndServe
var logFatal = log.Fatal

// newHandler wires the greeting service with the HTTP adapter.
func newHandler() http.Handler {
	greeter := service.NewGreetingService()
	return httpadapter.NewHandler(greeter)
}

func main() {
	log.Println("starting server on :8080")
	if err := listenAndServe(":8080", newHandler()); err != nil {
		logFatal(err)
	}
}
