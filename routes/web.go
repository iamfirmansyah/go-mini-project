package routes

import (
	"go-pos/controllers"

	"github.com/gorilla/mux"
)

func Public(router *mux.Router) {
	router.HandleFunc("/products", controllers.GetAll).Methods("GET")
	router.HandleFunc("/product", controllers.GetByID).Methods("GET")
	router.HandleFunc("/product", controllers.Create).Methods("POST")
	router.HandleFunc("/product", controllers.Update).Methods("PUT")
	router.HandleFunc("/product", controllers.Delete).Methods("DELETE")
}
