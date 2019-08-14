package app

import (
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

/*SetupRouter Setup Routers for the APIs. */
func (app *App) SetupRouter() {

	app.Router.
		Methods("POST").
		Path("/{table}").
		HandlerFunc(app.postFunction)

	app.Router.
		Methods("PUT").
		Path("/{table}").
		HandlerFunc(app.putFunction)

	app.Router.
		Methods("GET").
		Path("/{table}/findByStatus/{status}").
		HandlerFunc(app.getStatusFunction)

	app.Router.
		Methods("GET").
		Path("/{table}/{id}").
		HandlerFunc(app.createHandlerFunc())

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
func (app *App) createHandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	// var w http.ResponseWriter
	// var r *http.Request
	// vars := mux.Vars(r)
	// id, ok := vars["id"]
	// table, ok := vars["table"]

	// if !ok {
	// 	log.Panic("No ID in the path", ok)
	// }

	// log.Println("Table is:", table)
	// log.Println("id is:", id)
	return app.getFunction

	// switch table {
	// case "pet":
	// 	return app.getFunction
	// case "store":
	// 	return nil //app.getFunction
	// default:
	// 	return app.getFunction
	// }
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
func (app *App) getStatusFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, ok := vars["status"]
	if !ok {
		log.Fatal("No status in the path")
	}

	// Query all pets having status, append all pet to pets []Pet array;
	pet := &Pet{}
	var pets []Pet

	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
	var tmpPhotoUrls string

	// Use the tmpTags string to store the &pet.Tags string templetely, then Decode to []tag format.
	var tmpTags string

	// Using for counting how many records fetched
	count := 0
	rows, err := app.Database.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE status = ?", status)

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
			fmt.Println("tmpPhotoUrls from DB is: %s", tmpPhotoUrls)
			fmt.Println("tmpTags from DB is: %s", tmpTags)

			// Defining the photoUrlsJSON for Decode the string to []photourl
			var photoUrlsJSON []photourl
			decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
			errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)

			// Testing Decoded values of photoUrlsJSON
			fmt.Println("After Decode tmpPhotoUrls err is %s, array is %s", errPhotoUrlsJSON, photoUrlsJSON)

			// Defining the tagsJSON for Decode the string to []photourl
			var tagsJSON []tags
			decodeTmpTags := json.NewDecoder(strings.NewReader(tmpTags))
			errTagsJSON := decodeTmpTags.Decode(&tagsJSON)

			// Testing Decoded values of tagsJSON
			fmt.Println("After Decode tmpTags err is %s, array is %s", errTagsJSON, tagsJSON)

			if err != nil {
				log.Fatal("Database SELECT failed", err)
			}

			// Append *pet to pets for the final output
			pets = append(pets, *pet)
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
* Handle the get method to query record and pass the JSON data to front end.
 */
func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	table, ok := vars["table"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	log.Println("Table is:", table)
	log.Println("id is:", id)

	// Query all pets having status, append all pet to pets []Pet array;
	pet := &Pet{}
	var pets []Pet

	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
	var tmpPhotoUrls string

	// Use the tmpTags string to store the &pet.Tags string templetely, then Decode to []tag format.
	var tmpTags string

	// Using for counting how many records fetched
	count := 0
	rows, err := app.Database.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE id = ?", id)

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
			fmt.Println("tmpPhotoUrls from DB is: %s", tmpPhotoUrls)
			fmt.Println("tmpTags from DB is: %s", tmpTags)

			// Defining the photoUrlsJSON for Decode the string to []photourl
			var photoUrlsJSON []photourl
			decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
			errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)

			// Testing Decoded values of photoUrlsJSON
			fmt.Println("After Decode tmpPhotoUrls err is %s, array is %s", errPhotoUrlsJSON, photoUrlsJSON)

			// Defining the tagsJSON for Decode the string to []photourl
			var tagsJSON []tags
			decodeTmpTags := json.NewDecoder(strings.NewReader(tmpTags))
			errTagsJSON := decodeTmpTags.Decode(&tagsJSON)

			// Testing Decoded values of tagsJSON
			fmt.Println("After Decode tmpTags err is %s, array is %s", errTagsJSON, tagsJSON)

			if err != nil {
				log.Fatal("Database SELECT failed", err)
			}

			// Append *pet to pets for the final output
			pets = append(pets, *pet)
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
* Handle the post method to update record by petID.
 */
func (app *App) postUpdateFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No id in the path")
	}

	pet := &Pet{}
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

	pet := &Pet{}
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

	pet := &Pet{}
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
