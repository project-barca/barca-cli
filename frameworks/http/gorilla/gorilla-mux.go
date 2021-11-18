package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/v1/api/properties", getProperties).Methods("GET")
	route.HandleFunc("/v1/api/properties/{id}", getProperty).Methods("GET")
	route.HandleFunc("/v1/api/properties", createProperty).Methods("POST")
	route.HandleFunc("/v1/api/properties/{id}", updateProperty).Methods("PUT")
	route.HandleFunc("/v1/api/properties/{id}", deleteProperty).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8880", route))
}
