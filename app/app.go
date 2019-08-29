package app

import (
	"github.com/gorilla/mux"
)

// App struct for defining the Router
type App struct {
	Router *mux.Router
}

// SetupRouter Setup Routers for the APIs.
// Redirect all the requests to different Params HandlerFunc.
func (app *App) SetupRouter() {
	// Handler one param Methods
	app.Router.
		Methods("POST").
		Path("/{param1}").
		HandlerFunc(app.oneParamsHandlerFunc)

	app.Router.
		Methods("PUT").
		Path("/{param1}").
		HandlerFunc(app.oneParamsHandlerFunc)

	// Handler two params  Method
	app.Router.
		Methods("GET").
		Path("/{param1}/{param2}").
		HandlerFunc(app.twoParamsHandlerFunc)

	app.Router.
		Methods("POST").
		Path("/{param1}/{param2}").
		HandlerFunc(app.twoParamsHandlerFunc)

	app.Router.
		Methods("DELETE").
		Path("/{param1}/{param2}").
		HandlerFunc(app.twoParamsHandlerFunc)

	// Handler three params  Method
	app.Router.
		Methods("GET").
		Path("/{param1}/{param2}/{param3}").
		HandlerFunc(app.threeParamsHandlerFunc)

	app.Router.
		Methods("POST").
		Path("/{param1}/{param2}/{param3}").
		HandlerFunc(app.threeParamsHandlerFunc)
}
