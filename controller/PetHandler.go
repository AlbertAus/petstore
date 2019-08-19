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
	"strconv"
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

/*PetPostFunction handling the Post method to add new record to the database. */
func PetPostFunction(w http.ResponseWriter, r *http.Request) {
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

/*PetPutFunction handling the Put method to update the record to the database. */
func PetPutFunction(w http.ResponseWriter, r *http.Request) {
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

/*PetGetStatusFunction handling the get findByStatus method to query record by Status and pass the JSON data to front end. */
func PetGetStatusFunction(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// param3, ok := vars["param3"]
	param3, ok := r.URL.Query()["status"]
	if !ok || len(param3[0]) < 1 {
		// log.Panic("No status in the path")
		http.Error(w, "Invalid status value", 405)
		return
	}

	var checkParam3 models.Status
	//status for SQL Query use.
	var statusQry string
	for i := 0; i < len(param3); i++ {
		// var checkParam3 models.Status
		checkParam3 = models.Status(param3[i])

		// Add the param3[i] to statusQry string for SQL statement
		statusQry = statusQry + "," + "'" + param3[i] + "'"

		// Checking the input status ISValid
		if !checkParam3.IsValid() {
			fmt.Println(checkParam3)
			http.Error(w, "Invalid status value", 405)
			return
		}
	}

	// Remove the first "," in statusQry
	statusQry = statusQry[1:]
	fmt.Println("statusQry string is:" + statusQry)
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

	rows, err := DB.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE status IN (" + statusQry + ")")

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

/*PetPostUpdateFunction handling the post method to update record by petID. */
func PetPostUpdateFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("********* Entering the controller PetPostUpdateFunction(w,r) *********")
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
		var checkStatus models.Status
		checkStatus = models.Status(status)

		if !checkStatus.IsValid() {
			fmt.Println(checkStatus)
			http.Error(w, "Invalid input", 405)
			return
		}

		// Check the input Pet id's pet is Exists or not.
		var exists bool
		row := DB.QueryRow("SELECT EXISTS(SELECT * FROM `pet` WHERE id = ?)", id)
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

/*PetDeleteFunction handling the delete method to delete record by petID. */
func PetDeleteFunction(w http.ResponseWriter, r *http.Request) {
	log.Println("********* Entering the controller PetDeleteFunction(w,r) *********")
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

// petNotFound function handling the Pet Not Found message
func petNotFound(param2 string, w http.ResponseWriter) {
	// Output Pet Not Found message.
	var errMessage models.PetNotFound
	errMessage.Code, _ = strconv.ParseInt(param2, 10, 64)
	errMessage.Type = "error"
	errMessage.Message = "Pet not found"

	// Write the response text to front end
	outputMessage, outputErr := json.Marshal(errMessage)
	if outputErr == nil {
		http.Error(w, string(outputMessage), 404)
	}
}
