package route

import (
	"../controller"
	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	//routes
	router.HandleFunc("/api/register", controller.Register).Methods("POST")
	router.HandleFunc("/tasks", controller.Setup).Methods("GET")
}
