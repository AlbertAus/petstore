package controller

import (
	db "PetStore/database"
	"PetStore/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

/*PetUploadImageFunction handling the post method to upload a pet's Image by petID. */
func PetUploadImageFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("********* Entering the controller PetUploadImageFunction(w,r) *********")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `file`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our uploadimages directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("uploadimages", "pet-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "successful operation\n")

	tempFilePath, errPath := filepath.Abs(tempFile.Name())

	if errPath != nil {
		fmt.Println("Get tempFile path failed", err)
	}

	fmt.Println("The file path is: ", tempFilePath)

	// Setup DB to create Database connection and defer to Close() the DB connection
	DB, err := db.CreateDatabase()
	if err != nil {
		log.Panic("Database connection error!")
	}
	defer DB.Close()

	vars := mux.Vars(r)
	param2, ok2 := vars["param2"]
	if !ok2 {
		http.Error(w, "Invalid input", 405)
		return
	}

	log.Println("param2 is:", param2)

	additionalMetadata := r.FormValue("additionalMetadata")
	fmt.Printf("additionalMetadata is: %v\n", additionalMetadata)

	var pet Pet
	var responseBody models.ResponseBody

	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
	var tmpPhotoUrls string
	if r.Method == "POST" {

		// Using for counting how many records fetched
		count := 0

		rows, err := DB.Query("SELECT id, photoUrls FROM `pet` WHERE id = ?", param2)

		// Loop the records, do all Types convertion, append the result to pets if no error
		for rows.Next() {
			// Switch the err and print the outcomes.
			switch err := rows.Scan(&pet.ID, &tmpPhotoUrls); err {
			case sql.ErrNoRows:
				petNotFound(param2, w)
				return

			case nil:
				// Testing tmpPhotoUrls values getting from the database
				fmt.Printf("tmpPhotoUrls from DB is: %s\n", tmpPhotoUrls)

				// Defining the photoUrlsJSON for Decode the string to []photourl
				var photoUrlsJSON []string
				decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
				errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)
				pet.PhotoUrls = photoUrlsJSON

				// append the uploaded file path to PhotoUrls
				pet.PhotoUrls = append(pet.PhotoUrls, tempFilePath)

				// Testing Decoded values of photoUrlsJSON
				fmt.Printf("After Decode tmpPhotoUrls err is %s\n, array is %s\n", errPhotoUrlsJSON, photoUrlsJSON)
				fmt.Printf("The new pet.PhotoUrls is %s\n", pet.PhotoUrls)

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

		log.Printf("There are %v records fetched!\n", count)
		log.Println("You fetched pets by id!")

		photourls, urlsErr := json.Marshal(pet.PhotoUrls)
		if urlsErr != nil {
			fmt.Println("Converting pet.Category failed.", urlsErr)
			http.Error(w, "Validation exception", 405)
			return
		}

		updatePhotoUrls, err := DB.Prepare("UPDATE `pet` SET photoUrls =? WHERE id = ?")
		if err != nil {
			log.Panic("Database UPDATE failed")
		}

		updatePhotoUrls.Exec(photourls, param2)
		fmt.Println("You updated a pet by: " + param2 + "!")
	}
	responseBody.Code = 200
	responseBody.Type = handler.Header["Content-Type"][0]
	responseBody.Message = "additionalMetadata:" + additionalMetadata + "\n" + " File uploaded to " + tempFile.Name()
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseBody); err != nil {
		panic(err)
	}
}
