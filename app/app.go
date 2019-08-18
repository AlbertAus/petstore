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

/**
*	oneParamsHandlerFunc use for calling different handlers by the Paths with one Parameters.
 */
func (app *App) oneParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param1, ok1 := vars["param1"]

	if !ok1 {
		log.Panic("No Pet in the path", ok1)
	}

	fmt.Println("param1 is:", param1)
	if r.Method == "POST" {
		switch param1 {
		case "pet":
			fmt.Println("Calling Pet POST handler function")
			Handler.PetPostFunction(w, r)
		case "store":
			fmt.Println("Calling store POST handler function")
		case "user":
			fmt.Println("Calling user POST handler function")
		default:
			fmt.Println("Wrong URL!")
		}
	}

	if r.Method == "PUT" {
		switch param1 {
		case "pet":
			fmt.Println("Calling Pet PUT handler function")
			Handler.PetPutFunction(w, r)
		case "store":
			fmt.Println("Calling store PUT handler function")
		case "user":
			fmt.Println("Calling user PUT handler function")
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

	// Handle two params GET METHODs
	if r.Method == "GET" {
		switch param1 {
		case "pet":
			if param2 == "findByStatus" {
				fmt.Println("Calling Pet findByStatus handler function")
				Handler.PetGetStatusFunction(w, r)
			} else {
				fmt.Println("Calling PetGetFunction handler function")
				Handler.PetGetFunction(w, r)
			}
		case "store":
			fmt.Println("Calling store handler function")
		case "user":
			fmt.Println("Calling user handler function")
		default:
			fmt.Println("Wrong URL!")
		}
	}

	// Handle two params POST METHODs
	if r.Method == "POST" {
		switch param1 {
		case "pet":
			fmt.Println("Calling PetPostUpdateFunction handler function")
			Handler.PetPostUpdateFunction(w, r)
		case "store":
			fmt.Println("Calling store handler function")
		case "user":
			fmt.Println("Calling user handler function")
		default:
			fmt.Println("Wrong URL!")
		}
	}

	// Handle two params DELETE METHODs
	if r.Method == "DELETE" {
		switch param1 {
		case "pet":
			fmt.Println("Calling PetDeleteFunction handler function")
			Handler.PetDeleteFunction(w, r)
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
*	threeParamsHandlerFunc use for calling different handlers by the Paths with three Parameters.
 */
func (app *App) threeParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	param1, ok1 := vars["param1"]
	param2, ok2 := vars["param2"]
	param3, ok3 := vars["param3"]

	if !ok1 {
		log.Panic("Table name error in the path", ok1)
	}

	if !ok2 {
		log.Panic("No ID in the path", ok2)
	}

	if !ok3 {
		log.Panic("Param3 err in the path", ok3)
	}

	fmt.Println("param1 is:", param1)
	fmt.Println("param2 is:", param2)
	fmt.Println("param3 is:", param3)

	// Handle three params GET METHODs
	if r.Method == "GET" {
		switch param1 {
		case "pet":
			if param2 == "findByStatus" {
				fmt.Println("Calling Pet findByStatus handler function")
				Handler.PetGetStatusFunction(w, r)
			}

		case "store":
			if param2 == "inventory" {
				fmt.Println("Calling store inventory handler function")
			} else if param2 == "order" {
				fmt.Println("Calling store order handler function")
			}

		case "user":
			if param2 == "createWithArray" {
				fmt.Println("Calling user createWithArray handler function")
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

	// Handle three params POST METHODs
	if r.Method == "POST" {
		switch param1 {
		case "pet":
			if param3 == "uploadImage" {
				fmt.Println("Calling Pet uploadImage handler function")
				Handler.PetUploadImageFunction(w, r)
			}

		case "store":
			if param2 == "inventory" {
				fmt.Println("Calling store inventory handler function")
			} else if param2 == "order" {
				fmt.Println("Calling store order handler function")
			}

		case "user":
			if param2 == "createWithArray" {
				fmt.Println("Calling user createWithArray handler function")
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
}
