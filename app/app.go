package app

import (
	pet "github.com/AlbertAus/petstore/controller/pet"
	"github.com/gorilla/mux"
)

// App struct for defining the Router
type App struct {
	Router *mux.Router
}

// SetupRouter Setup Routers for the APIs.
// Redirect all the requests to different Params HandlerFunc.
func (app *App) SetupRouter() {

	// Handle pet router
	app.Router.
		Methods("POST").
		Path("/pet").HandlerFunc(pet.Post)

	app.Router.
		Methods("PUT").
		Path("/pet").HandlerFunc(pet.Put)

	app.Router.
		Methods("GET").
		Path("/pet/{param2}").HandlerFunc(pet.Get)

	app.Router.
		Methods("POST").
		Path("/pet/{param2}").HandlerFunc(pet.PostUpdate)

	app.Router.
		Methods("DELETE").
		Path("/pet/{param2}").
		HandlerFunc(pet.Delete)

	app.Router.
		Methods("POST").
		Path("/pet/{param2}/uploadImage").
		HandlerFunc(pet.UploadImage)

}
