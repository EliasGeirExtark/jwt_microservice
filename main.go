package main

import (
	"fmt"
	"github.com/extark/jwt_microservice/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var config Config

func main() {
	err := initSettings()
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	// create new router with the base prefix /api
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/refresh", nil).Methods("POST")

	fmt.Println("Server started...")
	log.Panicln(http.ListenAndServe(config.PORT, router))
}
