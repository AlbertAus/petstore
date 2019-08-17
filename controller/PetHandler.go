package controller

import (
	db "PetStore/database"
	"PetStore/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Pet defined for Global usuage
type Pet models.Pet

//  pet, pets, append all pet to pets []Pet array
// var pet Pet
// var pets []Pet

/*PetGetFunction handling the get method to query record and pass the JSON data to front end. */
func PetGetFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("********* Entering the controller PetGetFunction(w,r) *********")
	vars := mux.Vars(r)
	param2, ok2 := vars["param2"]
	if !ok2 {
		http.Error(w, "Invalid ID supplied", 400)
		return
	}

	log.Println("param2 is:", param2)

	var pet Pet

	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
	var tmpPhotoUrls string

	// Use the tmpTags string to store the &pet.Tags string templetely, then Decode to []tag format.
	var tmpTags string

	// Using for counting how many records fetched
	count := 0

	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE id = ?", param2)

	// Loop the records, do all Types convertion, append the result to pets if no error
	for rows.Next() {
		// Switch the err and print the outcomes.
		switch err := rows.Scan(&pet.ID, &pet.Category.ID, &pet.Category.Name, &pet.Name, &tmpPhotoUrls, &tmpTags, &pet.Status); err {
		case sql.ErrNoRows:
			http.Error(w, "Pet not found", 404)
			return

		case nil:
			// Testing tmpPhotoUrls and tmpTags values getting from the database
			fmt.Printf("tmpPhotoUrls from DB is: %s\n", tmpPhotoUrls)
			fmt.Printf("tmpTags from DB is: %s\n", tmpTags)

			// Defining the photoUrlsJSON for Decode the string to []photourl
			var photoUrlsJSON []string
			decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
			errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)
			pet.PhotoUrls = photoUrlsJSON

			// Testing Decoded values of photoUrlsJSON
			fmt.Printf("After Decode tmpPhotoUrls err is %s\n, array is %s\n", errPhotoUrlsJSON, photoUrlsJSON)

			// Defining the tagsJSON for Decode the string to []photourl
			var tagsJSON []models.Tag
			decodeTmpTags := json.NewDecoder(strings.NewReader(tmpTags))
			errTagsJSON := decodeTmpTags.Decode(&tagsJSON)
			pet.Tags = tagsJSON

			// Testing Decoded values of tagsJSON
			fmt.Printf("After Decode tmpTags err is %v\n, array is %v\n", errTagsJSON, tagsJSON)

			if err != nil {
				// log.Panic("Database SELECT failed", err)
				http.Error(w, "Pet not found", 404)
				return
			}

			// Append *pet to pets for the final output, this is only for multipul records.
			// pets = append(pets, pet)
			count++
		default:
			log.Panic("Database SELECT failed: ", err)
		}
	}
	if err != nil {
		// log.Panic("Database SELECT failed", err)
		http.Error(w, "Pet not found", 404)
		return
	}

	// If no Pet found, then return 404 error and "Pet not found"
	if count == 0 {
		// log.Printf(w, "No rows were returned!\n")
		http.Error(w, "Pet not found", 404)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("There are %v records fetched!\n", count)
	log.Println("You fetched pets by id!")
	// w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}
}

/*PetPostFunction handling the Post method to add new record to the database. */
func PetPostFunction(w http.ResponseWriter, r *http.Request) {
	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, DBerr := db.CreateDatabase()
	if DBerr != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	var pet Pet

	if r.Method == "POST" {
		// Testing request POST body
		if r.Body == nil {
			http.Error(w, "Invalid input", 405)
			return
		}

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

		// id := r.FormValue("id")
		// name := r.FormValue("name")

		insForm, insPrepareErr := DB.Prepare("INSERT INTO `pet` ( id, category, name, photoUrls, tags, status ) VALUES (?,?,?,?,?,?)")
		if insPrepareErr != nil {
			// log.Panic("Database INSERT failed")
			http.Error(w, "Invalid input", 405)
			return
		}

		_, insErr := insForm.Exec(id, category, name, photourls, tags, status)
		if insErr != nil {
			// log.Panic("Insert into Pet failed.\n", insErr)
			http.Error(w, "Invalid input", 405)
			return
		}
		log.Println("You added a new pet, id is: " + string(id) + " | name is: " + name)
	}
	w.WriteHeader(http.StatusOK)
	if jsonErr := json.NewEncoder(w).Encode(pet); jsonErr != nil {
		panic(jsonErr)
	}

}

/*PutFunction handling the Put method to update the record to the database. */
func PutFunction(w http.ResponseWriter, r *http.Request) {
	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	if r.Method == "PUT" {
		id := r.FormValue("id")
		name := r.FormValue("name")

		insForm, err := DB.Prepare("UPDATE `pet` SET (id, name) VALUES (?,?)")
		if err != nil {
			log.Panic("Database INSERT failed")
		}

		insForm.Exec(id, name)
		log.Println("You updated a new pet, id is: " + id + " | name is: " + name)
	}
	w.WriteHeader(http.StatusOK)

}

/*PetGetStatusFunction handling the get findByStatus method to query record by Status and pass the JSON data to front end. */
func PetGetStatusFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param3, ok := vars["param3"]
	if !ok {
		log.Panic("No status in the path")
	}
	fmt.Println("param3 is: " + param3)
	// var checkParam3 models.Status
	checkParam3 := models.Status(param3)

	// Checking the input status ISValid
	if !checkParam3.IsValid() {
		fmt.Println(checkParam3)
		http.Error(w, "Invalid status value", 405)
		return
	}

	// // Query all pets having status, append all pet to pets []Pet array;
	// Clear the pets array
	var pet Pet
	var pets []Pet
	pets = nil

	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
	var tmpPhotoUrls string

	// Use the tmpTags string to store the &pet.Tags string templetely, then Decode to []tag format.
	var tmpTags string

	// Using for counting how many records fetched
	count := 0

	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE status IN (?)", param3)

	// Loop the records, do all Types convertion, append the result to pets if no error
	for rows.Next() {
		// Switch the err and print the outcomes.
		switch err := rows.Scan(&pet.ID, &pet.Category.ID, &pet.Category.Name, &pet.Name, &tmpPhotoUrls, &tmpTags, &pet.Status); err {
		case sql.ErrNoRows:
			// If no Pet found, then return 404 error and "Pet not found"
			http.Error(w, "Pet not found", 404)
			return
		case nil:
			// Testing tmpPhotoUrls and tmpTags values getting from the database
			fmt.Printf("tmpPhotoUrls from DB is: %s\n", tmpPhotoUrls)
			fmt.Printf("tmpTags from DB is: %s\n", tmpTags)

			// Defining the photoUrlsJSON for Decode the string to []photourl
			var photoUrlsJSON []string
			decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
			errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)
			pet.PhotoUrls = photoUrlsJSON

			// Testing Decoded values of photoUrlsJSON
			fmt.Printf("After Decode tmpPhotoUrls err is %s\n, array is %s\n", errPhotoUrlsJSON, photoUrlsJSON)

			// Defining the tagsJSON for Decode the string to []photourl
			var tagsJSON []models.Tag
			decodeTmpTags := json.NewDecoder(strings.NewReader(tmpTags))
			errTagsJSON := decodeTmpTags.Decode(&tagsJSON)
			pet.Tags = tagsJSON

			// Testing Decoded values of tagsJSON
			fmt.Printf("After Decode tmpTags err is %v\n, array is %v\n", errTagsJSON, tagsJSON)

			if err != nil {
				log.Panic("Database SELECT failed", err)
			}

			// Append *pet to pets for the final output
			pets = append(pets, pet)
			count++
		default:
			log.Panic("Database SELECT failed: ", err)
		}
	}
	if err != nil {
		log.Panic("Database SELECT failed", err)
	}

	// If no Pet found, then return 404 error and "Pet not found"
	if count == 0 {
		// log.Printf(w, "No rows were returned!\n")
		http.Error(w, "Pet not found", 404)
		return
	}

	// w.WriteHeader(http.StatusOK)
	log.Printf("There are %v records fetched!\n", count)
	log.Println("You fetched pets by Status!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pets); err != nil {
		panic(err)
	}
}

/*PostUpdateFunction handling the post method to update record by petID. */
func PostUpdateFunction(w http.ResponseWriter, r *http.Request) {

	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()
	var pet Pet
	if r.Method == "PUT" {
		id := r.FormValue("id")
		name := r.FormValue("name")

		insForm, err := DB.Prepare("UPDATE `pet` SET Pet=? WHERE id = ?")
		if err != nil {
			log.Panic("Database UPDATE failed")
		}

		insForm.Exec(name, id)
		log.Println("You updated a pet by: " + id + "!")
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}
}

/*DeleteFunction handling the delete method to delete record by petID. */
func DeleteFunction(w http.ResponseWriter, r *http.Request) {
	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	var pet Pet

	if r.Method == "PUT" {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			log.Panic("No id in the path")
		}

		insForm, err := DB.Prepare("DELETE FROM `pet` WHERE id =?")
		if err != nil {
			log.Panic("Database UPDATE failed")
		}

		insForm.Exec(id)
		log.Println("You updated a pet by: " + id + "!")
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}

}

/*UploadImageFunction handling the post method to upload a pet's Image by petID. */
func UploadImageFunction(w http.ResponseWriter, r *http.Request) {
	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	var pet Pet

	if r.Method == "PUT" {

		vars := mux.Vars(r)
		id, ok1 := vars["id"]
		image, ok2 := vars["image"]
		if !ok1 || !ok2 {
			log.Panic("No id in the path")
		}
		insForm, err := DB.Prepare("UPDATE `pet` SET image =? path WHERE id = ?")
		if err != nil {
			log.Panic("Database UPDATE failed")
		}

		insForm.Exec(image, id)
		log.Println("You updated a pet by: " + id + "!")
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}
}
