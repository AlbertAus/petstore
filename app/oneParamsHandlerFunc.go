package app

import (
	Handler "PetStore/controller/pet"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/**
*	oneParamsHandlerFunc use for calling different handlers by the Paths with one Parameters.
 */
func (app *App) oneParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Getting all the params from URL
	vars := mux.Vars(r)
	param1, ok1 := vars["param1"]

	if !ok1 {
		log.Panic("No Pet in the path", ok1)
	}

	// Checking the METHODS, param1 values, then call related function
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

	// Checking the METHODS, param1 values, then call related function
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
