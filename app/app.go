package app

import (
	"PetStore/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

/*App struct for defining the Router and Database */
type App struct {
	Router   *mux.Router
	Database *sql.DB
}

//Pet defined for Global usuage
type Pet models.Pet

//  pet, pets, append all pet to pets []Pet array
var pet Pet
var pets []Pet

func init() {
	// pet = models.Pet{}
}

/*SetupRouter Setup Routers for the APIs. */
func (app *App) SetupRouter() {

	// Handler two params GET Method
	app.Router.
		Methods("GET").
		Path("/{param1}/{param2}").
		HandlerFunc(app.twoParamsGetHandlerFunc)

	// Handler three params GET Method
	app.Router.
		Methods("GET").
		Path("/{param1}/{param2}/{param3}").
		HandlerFunc(app.threeParamsGetHandlerFunc)

	app.Router.
		Methods("POST").
		Path("/{table}").
		HandlerFunc(app.postFunction)

	app.Router.
		Methods("PUT").
		Path("/{table}").
		HandlerFunc(app.putFunction)

	app.Router.
		Methods("POST").
		Path("/{table}/{id}").
		HandlerFunc(app.postUpdateFunction)

	app.Router.
		Methods("DELETE").
		Path("/{table}/{id}").
		HandlerFunc(app.deleteFunction)

	app.Router.
		Methods("POST").
		Path("/{table}/{id}/uploadImage").
		HandlerFunc(app.uploadImageFunction)
}

/**
*	createHandlerFunc use for calling different handlers by the Path's Parameters
 */
func (app *App) twoParamsGetHandlerFunc(w http.ResponseWriter, r *http.Request) {

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
		app.petGetFunction(w, r)
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
*	createHandlerFunc use for calling different handlers by the Path's Parameters.
 */
func (app *App) threeParamsGetHandlerFunc(w http.ResponseWriter, r *http.Request) {

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
			app.petGetStatusFunction(w, r)
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

/**
* Handle the get method to query record and pass the JSON data to front end.
 */
func (app *App) petGetFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// param1, ok1 := vars["param1"]
	param2, ok2 := vars["param2"]
	// if !ok1 {
	// 	log.Fatal("No pet in the path")
	// }
	if !ok2 {
		log.Fatal("No ID in the path")
	}
	// log.Println("param1 is:", param1)
	log.Println("param2 is:", param2)

	// // Query all pets having status, append all pet to pets []Pet array;
	// pet := &models.Pet{}
	// var pets []models.Pet
	// Clear the pets array
	pets = nil

	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
	var tmpPhotoUrls string

	// Use the tmpTags string to store the &pet.Tags string templetely, then Decode to []tag format.
	var tmpTags string

	// Using for counting how many records fetched
	count := 0
	rows, err := app.Database.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE id = ?", param2)

	// Loop the records, do all Types convertion, append the result to pets if no error
	for rows.Next() {
		// Switch the err and print the outcomes.
		switch err := rows.Scan(&pet.ID, &pet.Category.ID, &pet.Category.Name, &pet.Name, &tmpPhotoUrls, &tmpTags, &pet.Status); err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Welcome to PetStore!\n")
			if err := json.NewEncoder(w).Encode("No rows were returned!"); err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "No rows were returned!\n")
		case nil:
			// Testing tmpPhotoUrls and tmpTags values getting from the database
			fmt.Printf("tmpPhotoUrls from DB is: %s", tmpPhotoUrls)
			fmt.Printf("tmpTags from DB is: %s", tmpTags)

			// Defining the photoUrlsJSON for Decode the string to []photourl
			var photoUrlsJSON []models.Photourl
			decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
			errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)

			// Testing Decoded values of photoUrlsJSON
			fmt.Printf("After Decode tmpPhotoUrls err is %s, array is %s", errPhotoUrlsJSON, photoUrlsJSON)

			// Defining the tagsJSON for Decode the string to []photourl
			var tagsJSON []models.Tags
			decodeTmpTags := json.NewDecoder(strings.NewReader(tmpTags))
			errTagsJSON := decodeTmpTags.Decode(&tagsJSON)

			// Testing Decoded values of tagsJSON
			fmt.Printf("After Decode tmpTags err is %s, array is %s", errTagsJSON, tagsJSON)

			if err != nil {
				log.Fatal("Database SELECT failed", err)
			}

			// Append *pet to pets for the final output
			pets = append(pets, pet)
			count++
		default:
			log.Panic("Database SELECT failed: ", err)
		}
	}
	if err != nil {
		log.Fatal("Database SELECT failed", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "There are %v records fetched!\n", count)
	log.Println("You fetched pets by id!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pets); err != nil {
		panic(err)
	}
}

/**
* Handle the Post method to add new record to the database.
 */
func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `pet` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	log.Println("You added a new pet!")
	w.WriteHeader(http.StatusOK)
}

/**
* Handle the Put method to update the record to the database.
 */
func (app *App) putFunction(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("UPDATE `pet` SET (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database UPDATE failed")
	}

	log.Println("You updated a pet!")
	w.WriteHeader(http.StatusOK)
}

/**
* Handle the get findByStatus method to query record by Status and pass the JSON data to front end.
 */
func (app *App) petGetStatusFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param3, ok := vars["param3"]
	if !ok {
		log.Fatal("No status in the path")
	}

	// // Query all pets having status, append all pet to pets []Pet array;
	// pet := &models.Pet{}
	// var pets []models.Pet
	// Clear the pets array
	pets = nil

	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
	var tmpPhotoUrls string

	// Use the tmpTags string to store the &pet.Tags string templetely, then Decode to []tag format.
	var tmpTags string

	// Using for counting how many records fetched
	count := 0
	rows, err := app.Database.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE status IN (?)", param3)

	// Loop the records, do all Types convertion, append the result to pets if no error
	for rows.Next() {
		// Switch the err and print the outcomes.
		switch err := rows.Scan(&pet.ID, &pet.Category.ID, &pet.Category.Name, &pet.Name, &tmpPhotoUrls, &tmpTags, &pet.Status); err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Welcome to PetStore!\n")
			if err := json.NewEncoder(w).Encode("No rows were returned!"); err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "No rows were returned!\n")
		case nil:
			// Testing tmpPhotoUrls and tmpTags values getting from the database
			fmt.Printf("tmpPhotoUrls from DB is: %s", tmpPhotoUrls)
			fmt.Printf("tmpTags from DB is: %s", tmpTags)

			// Defining the photoUrlsJSON for Decode the string to []photourl
			var photoUrlsJSON []models.Photourl
			decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
			errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)

			// Testing Decoded values of photoUrlsJSON
			fmt.Printf("After Decode tmpPhotoUrls err is %s, array is %s", errPhotoUrlsJSON, photoUrlsJSON)

			// Defining the tagsJSON for Decode the string to []photourl
			var tagsJSON []models.Tags
			decodeTmpTags := json.NewDecoder(strings.NewReader(tmpTags))
			errTagsJSON := decodeTmpTags.Decode(&tagsJSON)

			// Testing Decoded values of tagsJSON
			fmt.Printf("After Decode tmpTags err is %s, array is %s", errTagsJSON, tagsJSON)

			if err != nil {
				log.Fatal("Database SELECT failed", err)
			}

			// Append *pet to pets for the final output
			pets = append(pets, pet)
			count++
		default:
			log.Panic("Database SELECT failed: ", err)
		}
	}
	if err != nil {
		log.Fatal("Database SELECT failed", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "There are %v records fetched!\n", count)
	log.Println("You fetched pets by Status!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pets); err != nil {
		panic(err)
	}
}

/**
* Handle the post method to update record by petID.
 */
func (app *App) postUpdateFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No id in the path")
	}

	// pet := &models.Pet{}
	_, err := app.Database.Exec("UPDATE `pet` SET Pet=? WHERE id = ?", id, id)
	if err != nil {
		log.Fatal("Database UPDATE failed")
	}

	log.Println("You updated a pet by: " + id + "!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}
}

/**
* Handle the delete method to delete record by petID.
 */
func (app *App) deleteFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No id in the path")
	}

	// pet := &models.Pet{}
	_, err := app.Database.Exec("DELETE FROM `pet` WHERE id =?", id)
	if err != nil {
		log.Fatal("Database DELETE failed")
	}

	log.Println("You deleted a pet by: " + id + "!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}
}

/**
* Handle the post method to upload a pet's Image by petID.
 */
func (app *App) uploadImageFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	image, ok := vars["image"]
	if !ok {
		log.Fatal("No id in the path")
	}

	// pet := &models.Pet{}
	_, err := app.Database.Exec("UPDATE `pet` SET image =? path WHERE id = ?", image, id)
	if err != nil {
		log.Fatal("Database UPDATE failed")
	}

	log.Println("You uploaded a pet image!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}
}
