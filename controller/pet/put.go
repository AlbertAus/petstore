package pet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/AlbertAus/petstore/database"
)

// Put handling the Put method to update the record to the database.
func Put(w http.ResponseWriter, r *http.Request) {
	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, DBerr := db.CreateDatabase()
	if DBerr != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	// Defining the pet for Query use.
	var pet Pet

	if r.Method == "PUT" {
		// Testing request POST body
		if r.Body == nil {
			http.Error(w, "Validation exception", 405)
			return
		}

		// Decode r.Body to pet struct format
		inputErr := json.NewDecoder(r.Body).Decode(&pet)
		if inputErr != nil {
			http.Error(w, "Validation exception", 405)
			return
		}
		fmt.Printf("PUT going to update pet is: %v\n", pet)
		id := pet.ID
		name := pet.Name
		category, categoryErr := json.Marshal(pet.Category)
		if categoryErr != nil {
			fmt.Println("Converting pet.Category failed.", categoryErr)
			http.Error(w, "Validation exception", 405)
			return
		}
		photourls, urlsErr := json.Marshal(pet.PhotoUrls)
		if urlsErr != nil {
			fmt.Println("Converting pet.Category failed.", urlsErr)
			http.Error(w, "Validation exception", 405)
			return
		}
		tags, tagsErr := json.Marshal(pet.Tags)
		if tagsErr != nil {
			fmt.Println("Converting pet.Category failed.", tagsErr)
			http.Error(w, "Validation exception", 405)
			return
		}
		status := pet.Status

		fmt.Println(pet.Status.IsValid())
		// Checking the input status ISValid
		if !pet.Status.IsValid() {
			http.Error(w, "Validation exception", 405)
			return
		}

		// Check the input Pet id's pet is Exists or not.
		var exists bool
		row := DB.QueryRow("SELECT EXISTS(SELECT * FROM `pet` WHERE id = ?)", id)
		if existsErr := row.Scan(&exists); existsErr != nil {
			fmt.Println(existsErr)
		}

		// UPDATE the pet if Exists, other wise return Invalid ID supplied
		if exists {
			fmt.Println("Exists!")
			updateSQL, updateSQLPrepareErr := DB.Prepare("UPDATE `pet` SET id = ?, category = ?, name = ?, photoUrls = ?, tags = ?, status = ? WHERE id = ?")
			if updateSQLPrepareErr != nil {
				fmt.Println("Database UPDATE failed", updateSQLPrepareErr)
				http.Error(w, "Validation exception", 405)
				return
			}

			updateRes, updateSQLErr := updateSQL.Exec(id, category, name, photourls, tags, status, id)
			if updateSQLErr != nil {
				fmt.Println("UPDATE Pet failed.\n", updateSQLErr)
				http.Error(w, "Validation exception", 405)
				return
			}

			// Count Affected Rows, if countUpdated == 0, means Pet not found.
			countUpdated, err2 := updateRes.RowsAffected()

			if err2 != nil {
				fmt.Println(err2.Error())
			} else {
				fmt.Printf("%v record have been updated!\n", countUpdated)
				if countUpdated == 0 {
					http.Error(w, "Validation exception", 405)
					return
				}
			}

			fmt.Printf("You updated a pet, id is: %v, name is: %s \n", id, name)
		} else if !exists {
			fmt.Println("not exists!")
			http.Error(w, "Invalid ID supplied", 400)
			return
		}
	}

	// Write the response to front end
	w.WriteHeader(http.StatusOK)
	if jsonErr := json.NewEncoder(w).Encode(pet); jsonErr != nil {
		panic(jsonErr)
	}

}
