package pet

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/AlbertAus/petstore/database"

	"github.com/gorilla/mux"
)

// Delete handling the delete method to delete record by petID.
func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("******* Entering the controller PetDeleteFunction(w,r) *********")
	vars := mux.Vars(r)
	param2, ok2 := vars["param2"]
	if !ok2 {
		http.Error(w, "Invalid input", 405)
		return
	}

	log.Println("param2 is:", param2)

	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	if r.Method == "DELETE" {
		// Testing request POST body
		id := param2

		// Check the input Pet id's pet is Exists or not.
		var exists bool
		row := DB.QueryRow("SELECT EXISTS(SELECT * FROM `pet` WHERE id = ?)", id)
		if existsErr := row.Scan(&exists); existsErr != nil {
			fmt.Println(existsErr)
		}
		if exists {
			fmt.Println("Exists!")
			updateSQL, updateSQLPrepareErr := DB.Prepare("DELETE FROM `pet` WHERE id =?")
			if updateSQLPrepareErr != nil {
				fmt.Println("Database DELETE failed", updateSQLPrepareErr)
				http.Error(w, "Invalid input", 405)
				return
			}

			updateRes, updateSQLErr := updateSQL.Exec(id)
			if updateSQLErr != nil {
				fmt.Println("DELETE Pet failed.\n", updateSQLErr)
				http.Error(w, "Invalid input", 405)
				return
			}

			// Count Affected Rows, if countUpdated == 0, means Pet not found.
			countUpdated, err2 := updateRes.RowsAffected()

			if err2 != nil {
				fmt.Println(err2.Error())
			} else {
				fmt.Printf("%v record have been DELETE!\n", countUpdated)
				if countUpdated == 0 {
					http.Error(w, "Invalid input", 405)
					return
				}
			}
			fmt.Printf("You DELETE a pet, id is: %v\n", id)
		} else if !exists {
			fmt.Println("not exists!")
			http.Error(w, "Invalid input", 405)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
