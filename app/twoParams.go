package app

import (
	"fmt"
	"log"
	"net/http"

	pet "github.com/AlbertAus/petstore/controller/pet"

	"github.com/gorilla/mux"
)

// twoParamsHandlerFunc use for calling different handlers by the Paths with two Parameters.
func (app *App) twoParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Getting all the params from URL
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
				fmt.Println("Calling Pet findByStatus handler function 2")
				pet.GetStatus(w, r)
			} else {
				fmt.Println("Calling PetGetFunction handler function")
				pet.Get(w, r)
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
			pet.PostUpdate(w, r)
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
			pet.Delete(w, r)
		case "store":
			fmt.Println("Calling store handler function")
		case "user":
			fmt.Println("Calling user handler function")
		default:
			fmt.Println("Wrong URL!")
		}
	}
}
