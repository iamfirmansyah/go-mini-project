package main

import (
	"fmt"
	"go-pos/app"
	"go-pos/config"
	"go-pos/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// * Load Configuration
	config.LoadAppConfig()

	// * Initialize Database
	app.ConnectMysql(config.AppConfig.MYSQL_CONNECTION)

	// * Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	s := router.PathPrefix("/api").Subrouter()

	// * Register Routes Public
	routes.Product(s.PathPrefix("/product").Subrouter())

	// * Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.PORT), router))

}
