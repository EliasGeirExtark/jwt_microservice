package main

import (
	"fmt"
	"github.com/extark/jwt_microservice/controller"
	"github.com/extark/jwt_microservice/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	err := utils.InitSettings()
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	// create new router with the base prefix /api
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/refresh", controller.Refresh).Methods("POST")

	fmt.Println("Server started...")
	log.Panicln(http.ListenAndServe(utils.Cfg.PORT, router))
}
