package main

import (
	"github.com/italosm/rest-api/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.CreateRoutes()
	go log.Fatal(http.ListenAndServe(":8080", router))
}