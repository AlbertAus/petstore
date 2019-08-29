package pet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/AlbertAus/petstore/database"
)

// Post handling the Post method to add new record to the database.
func Post(w http.ResponseWriter, r *http.Request) {
	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, DBerr := db.CreateDatabase()
	if DBerr != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	// Defining the pet for Query use.
	var pet Pet

	if r.Method == "POST" {
		// Testing request POST body
		if r.Body == nil {
			http.Error(w, "Invalid input", 405)
			return
		}

		// Decode r.Body to pet struct format
		inputErr := json.NewDecoder(r.Body).Decode(&pet)
		if inputErr != nil {
			http.Error(w, "Invalid input", 405)
			return
		}
		fmt.Printf("Post pet is: %v\n", pet)

		id := pet.ID
		name := pet.Name
		category, categoryErr := json.Marshal(pet.Category)

		if categoryErr != nil {
			// log.Panic("Converting pet.Category failed.", categoryErr)
			http.Error(w, "Invalid input", 405)
			return
		}

		photourls, urlsErr := json.Marshal(pet.PhotoUrls)
		if urlsErr != nil {
			// log.Panic("Converting pet.Category failed.", urlsErr)
			http.Error(w, "Invalid input", 405)
			return
		}

		tags, tagsErr := json.Marshal(pet.Tags)
		if tagsErr != nil {
			// log.Panic("Converting pet.Category failed.", tagsErr)
			http.Error(w, "Invalid input", 405)
			return
		}
		status := pet.Status

		// fmt.Println(pet.Status.IsValid())

		// Checking the input status ISValid
		if !pet.Status.IsValid() {
			http.Error(w, "invalid input", 405)
			return
		}

		// INSERT INTO pet table with values.
		insForm, insPrepareErr := DB.Prepare("INSERT INTO `pet` ( id, category, name, photoUrls, tags, status ) VALUES (?,?,?,?,?,?)")
		if insPrepareErr != nil {
			http.Error(w, "Invalid input", 405)
			return
		}

		_, insErr := insForm.Exec(id, category, name, photourls, tags, status)
		if insErr != nil {
			// log.Panic("Insert into Pet failed.\n", insErr)
			http.Error(w, "Invalid input", 405)
			return
		}
		fmt.Printf("You inserted a pet, id is: %v, name is: %s \n", id, name)
	}

	// Write the response to front end.
	w.WriteHeader(http.StatusOK)
	if jsonErr := json.NewEncoder(w).Encode(pet); jsonErr != nil {
		panic(jsonErr)
	}

}
