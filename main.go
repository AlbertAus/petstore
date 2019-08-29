package main

import (
	"log"
	"net/http"

	"github.com/AlbertAus/petstore/app"

	"github.com/gorilla/mux"
)

// APP enter point
func main() {

	// Define the app, putting Router to the app
	app := &app.App{
		Router: mux.NewRouter().StrictSlash(true),
		// Database: database,
	}

	// Call the app SetupRouter function, handling all METHODS
	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
