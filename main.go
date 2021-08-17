package main

import (
	"Mobilebackend/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {


	log.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/getHeadlines", handlers.GetHeadlines).Methods("POST")
	log.Println("Service active and listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}