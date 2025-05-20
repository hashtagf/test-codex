package main

import (
	"log"
	"net/http"

	httpadapter "hexademo/internal/adapter/http"
	"hexademo/internal/service"
)

var (
	listenAndServe = http.ListenAndServe
	logFatal       = log.Fatal
)

func main() {
	greeter := service.NewGreetingService()
	handler := httpadapter.NewHandler(greeter)

	log.Println("starting server on :8080")
	if err := listenAndServe(":8080", handler); err != nil {
		logFatal(err)
	}
}
