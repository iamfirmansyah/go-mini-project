package routes

import (
	"go-pos/controllers"

	"github.com/gorilla/mux"
)

func Product(router *mux.Router) {
	router.HandleFunc("", controllers.GetAll).Methods("GET")
	router.HandleFunc("", controllers.GetByID).Methods("PATCH")
	router.HandleFunc("", controllers.Create).Methods("POST")
	router.HandleFunc("", controllers.Update).Methods("PUT")
	router.HandleFunc("", controllers.Delete).Methods("DELETE")
}
