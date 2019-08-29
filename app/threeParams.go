package app

import (
	"fmt"
	"log"
	"net/http"

	pet "github.com/AlbertAus/petstore/controller/pet"

	"github.com/gorilla/mux"
)

// threeParamsHandlerFunc use for calling different handlers by the Paths with three Parameters.
func (app *App) threeParamsHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Getting all the params from URL
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

	// Checking the METHODS, params values, then call related function
	// Handle three params GET METHODs
	if r.Method == "GET" {
		switch param1 {
		case "pet":
			// if param2 == "findByStatus" {
			// 	fmt.Println("Calling Pet findByStatus handler function 3")
			// 	handler.PetGetStatusFunction(w, r)
			// }

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

	// Checking the METHODS, params values, then call related function
	// Handle three params POST METHODs
	if r.Method == "POST" {
		switch param1 {
		case "pet":
			if param3 == "uploadImage" {
				fmt.Println("Calling Pet uploadImage handler function")
				pet.UploadImage(w, r)
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
