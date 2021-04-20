package main

import (
	"./database"
	"./route"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	router := mux.NewRouter().StrictSlash(true)

	route.Routes(router)
	//start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

