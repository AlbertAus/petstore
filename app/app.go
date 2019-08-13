package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
		Path("/pet").
		HandlerFunc(app.postFunction)

	app.Router.
		Methods("PUT").
		Path("/pet").
		HandlerFunc(app.putFunction)

	app.Router.
		Methods("GET").
		Path("/pet/findByStatus/{status}").
		HandlerFunc(app.getStatusFunction)

	app.Router.
		Methods("GET").
		Path("/pet/{id}").
		HandlerFunc(app.getFunction)

	app.Router.
		Methods("POST").
		Path("/pet/{id}").
		HandlerFunc(app.postUpdateFunction)

	app.Router.
		Methods("DELETE").
		Path("/pet/{id}").
		HandlerFunc(app.deleteFunction)

	app.Router.
		Methods("POST").
		Path("/pet/{id}/uploadImage").
		HandlerFunc(app.uploadImageFunction)
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
	// Using for counting how many records fetched
	count := 0
	rows, err := app.Database.Query("SELECT * FROM `pet` WHERE status = ?", status)
	for rows.Next() {
		err := rows.Scan(&pet.ID, &pet.Category, &pet.Name, &pet.PhotoUrls, &pet.tags, &pet.Status)
		if err != nil {
			log.Fatal("Database SELECT failed", err)
		}

		pets = append(pets, *pet)
		count++
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
	if !ok {
		log.Fatal("No ID in the path")
	}

	log.Println("id is:", id)

	pet := &Pet{}
	sqlStatement := `select * from pet where id =?;`
	log.Println("sqlStatement is:", sqlStatement)
	row := app.Database.QueryRow(sqlStatement, id)
	switch err := row.Scan(&pet.ID, &pet.Category, &pet.Name, &pet.PhotoUrls, &pet.tags, &pet.Status); err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to my website!\n")
		if err := json.NewEncoder(w).Encode("No rows were returned!"); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "No rows were returned!\n")
	case nil:
		log.Println("You fetched a pet!", &pet.ID, &pet.Category, &pet.Name, &pet.Status)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(pet); err != nil {
			panic(err)
		}
	default:
		log.Panic("Database SELECT failed: ", err)
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
