package main

import (
	"log"
	"net/http"

	httpadapter "hexademo/internal/adapter/http"
	"hexademo/internal/service"
)

func main() {
	greeter := service.NewGreetingService()
	handler := httpadapter.NewHandler(greeter)

	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
