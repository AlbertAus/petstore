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

/*PetGetFunction handling the get method to query record and pass the JSON data to front end. */
func PetGetFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("********* Entering the controller PetGetFunction(w,r) *********")

	// Getting all the params from URL
	vars := mux.Vars(r)
	param2, ok2 := vars["param2"]
	if !ok2 {
		http.Error(w, "Invalid ID supplied", 400)
		return
	}

	log.Println("param2 is:", param2)

	// Defining the pet for Query use.
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
			petNotFound(param2, w)
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
				petNotFound(param2, w)
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
		petNotFound(param2, w)
		return
	}

	// If no Pet found, then return 404 error and "Pet not found"
	if count == 0 {
		petNotFound(param2, w)
		return
	}

	// Write the response to front end.
	w.WriteHeader(http.StatusOK)
	log.Printf("There are %v records fetched!\n", count)
	log.Println("You fetched pets by id!")
	// w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pet); err != nil {
		panic(err)
	}
}
