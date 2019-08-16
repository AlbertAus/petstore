package app

import (
	Handler "PetStore/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*App struct for defining the Router*/
type App struct {
	Router *mux.Router
}

/*SetupRouter Setup Routers for the APIs.
* Redirect all the requests to different Params HandlerFunc.
 */
func (app *App) SetupRouter() {

	// Handler two params GET Method
	app.Router.
		Methods("GET").
		Path("/{param1}/{param2}").
		HandlerFunc(app.twoParamsHandlerFunc)

	// Handler three params GET Method
	app.Router.
		Methods("GET").
		Path("/{param1}/{param2}/{param3}").
		HandlerFunc(app.threeParamsHandlerFunc)

	app.Router.
		Methods("POST").
		Path("/{param1}").
		HandlerFunc(app.oneParamsHandlerFunc)

	app.Router.
		Methods("PUT").
		Path("/{param1}").
		HandlerFunc(Handler.PutFunction)

	app.Router.
		Methods("POST").
		Path("/{param1}/{param2}").
		HandlerFunc(Handler.PostUpdateFunction)

	app.Router.
		Methods("DELETE").
		Path("/{param1}/{param2}").
		HandlerFunc(Handler.DeleteFunction)

	app.Router.
		Methods("POST").
		Path("/{param1}/{param2}/{param3}").
		HandlerFunc(Handler.UploadImageFunction)
}

/**
*	oneParamsHandlerFunc use for calling different handlers by the Paths with one Parameters.
 */
func (app *App) oneParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	param1, ok1 := vars["param1"]
	// param2, ok2 := vars["param2"]

	if !ok1 {
		log.Panic("No Pet in the path", ok1)
	}

	// if !ok2 {
	// 	log.Panic("No ID in the path", ok2)
	// }

	fmt.Println("param1 is:", param1)
	// fmt.Println("param2 is:", param2)
	// return app.getFunction
	if r.Method == "POST" {
		switch param1 {
		case "pet":
			fmt.Println("Calling Pet handler function")
			// app.petGetFunction(w, r)
			Handler.PetPostFunction(w, r)
			// Call PetHandler function
			// controller.PetGetFunction(w, r)
		case "store":
			fmt.Println("Calling store handler function")
		case "user":
			fmt.Println("Calling user handler function")
		default:
			fmt.Println("Wrong URL!")
		}
	}
}

/**
*	twoParamsHandlerFunc use for calling different handlers by the Paths with two Parameters.
 */
func (app *App) twoParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	param1, ok1 := vars["param1"]
	param2, ok2 := vars["param2"]

	if !ok1 {
		log.Panic("No Pet in the path", ok1)
	}

	if !ok2 {
		log.Panic("No ID in the path", ok2)
	}

	fmt.Println("param1 is:", param1)
	fmt.Println("param2 is:", param2)
	// return app.getFunction

	switch param1 {
	case "pet":
		fmt.Println("Calling Pet handler function")
		// app.petGetFunction(w, r)
		Handler.PetGetFunction(w, r)
		// Call PetHandler function
		// controller.PetGetFunction(w, r)
	case "store":
		fmt.Println("Calling store handler function")
	case "user":
		fmt.Println("Calling user handler function")
	default:
		fmt.Println("Wrong URL!")
	}
}

/**
*	threeParamsHandlerFunc use for calling different handlers by the Paths with three Parameters.
 */
func (app *App) threeParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	param1, ok := vars["param1"]
	param2, ok := vars["param2"]
	param3, ok := vars["param3"]

	if !ok {
		log.Panic("No ID in the path", ok)
	}

	log.Println("param1 is:", param1)
	log.Println("Param2 is:", param2)
	log.Println("param3 is:", param3)
	// return app.getFunction

	switch param1 {
	case "pet":
		if param2 == "findByStatus" {
			fmt.Println("Calling Pet findByStatus handler function")
			Handler.PetGetStatusFunction(w, r)
		}

		if param3 == "uploadImage" {
			log.Println("Calling Pet uploadImage handler function")
		}

	case "store":
		if param2 == "inventory" {
			fmt.Println("Calling store inventory handler function")
			// app.petGetStatusFunction(w, r)
		} else if param2 == "order" {
			fmt.Println("Calling store order handler function")
		}

	case "user":
		if param2 == "createWithArray" {
			fmt.Println("Calling user createWithArray handler function")
			// app.petGetStatusFunction(w, r)
		} else if param2 == "createWithList" {
			fmt.Println("Calling user createWithList handler function")
		} else if param2 == "login" {
			fmt.Println("Calling user login handler function")
		} else if param2 == "logout" {
			fmt.Println("Calling user logout handler function")
		}
	default:
		fmt.Println("Calling other handler function")
	}
}
