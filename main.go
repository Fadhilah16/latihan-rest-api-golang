package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-crud-golang/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterProductRoutes(r)
	routes.RegisterAuthRoutes(r)
	http.Handle("/", r)

	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
