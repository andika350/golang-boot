package main

import (
	
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

// @title Orders API

// @version 1.0

// @description This is a sample service for managing orders

// @termOfService http://swagger.io/terms/

// @contact.name API Support

// @contact.email soberkoder@swagger.io

// @license.name Apache 2.0

// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

// @BasePath /


func main() {

	router := mux.NewRouter()
	 
	log.Fatal(http.ListenAndServe(":8080", router))

}
