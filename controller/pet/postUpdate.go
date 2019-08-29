package pet

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/AlbertAus/petstore/database"
	"github.com/AlbertAus/petstore/model"

	"github.com/gorilla/mux"
)

// PostUpdate handling the post method to update record by petID.
func PostUpdate(w http.ResponseWriter, r *http.Request) {
	log.Println("********// Entering the controller PetPostUpdateFunction(w,r) *********")
	vars := mux.Vars(r)
	param2, ok2 := vars["param2"]
	if !ok2 {
		http.Error(w, "Invalid input", 405)
		return
	}

	fmt.Println("param2 is:", param2)

	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, DBerr := db.CreateDatabase()
	if DBerr != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	// var pet Pet

	if r.Method == "POST" {
		// Testing request POST body
		id := param2
		name := r.FormValue("name")
		status := r.FormValue("status")
		fmt.Printf("name is: %v, status is: %v\n", name, status)
		if name == "" || status == "" {
			http.Error(w, "Invalid input", 405)
			return
		}

		// Using checkStatus to check the input status ISValid
		var checkStatus model.Status
		checkStatus = model.Status(status)

		if !checkStatus.IsValid() {
			fmt.Println(checkStatus)
			http.Error(w, "Invalid input", 405)
			return
		}

		// Check the input Pet id's pet is Exists or not.
		var exists bool
		row := DB.QueryRow("SELECT EXISTS(SELECT // FROM `pet` WHERE id = ?)", id)
		if existsErr := row.Scan(&exists); existsErr != nil {
			fmt.Println(existsErr)
		}
		if exists {
			fmt.Println("Exists!")
			updateSQL, updateSQLPrepareErr := DB.Prepare("UPDATE `pet` SET name = ?,status = ? WHERE id = ?")
			if updateSQLPrepareErr != nil {
				fmt.Println("Database UPDATE failed", updateSQLPrepareErr)
				http.Error(w, "Invalid input", 405)
				return
			}

			updateRes, updateSQLErr := updateSQL.Exec(name, status, id)
			if updateSQLErr != nil {
				fmt.Println("Update Pet failed.\n", updateSQLErr)
				http.Error(w, "Invalid input", 405)
				return
			}

			// Count Affected Rows, if countUpdated == 0, means Pet not found.
			countUpdated, err2 := updateRes.RowsAffected()

			if err2 != nil {
				fmt.Println(err2.Error())
			} else {
				fmt.Printf("%v record have been updated!\n", countUpdated)
				if countUpdated == 0 {
					http.Error(w, "Invalid input", 405)
					return
				}
			}

			fmt.Printf("You updated a pet, id is: %v, name is: %s, status is: %s \n", id, name, status)
		} else if !exists {
			fmt.Println("not exists!")
			http.Error(w, "Invalid input", 405)
			return
		}
	}
	w.WriteHeader(http.StatusOK)

}
