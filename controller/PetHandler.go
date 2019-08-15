package controller

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"github.com/gorilla/mux"
// )

// // func main() {
// // 	app := &app.App{
// // 		Router:   mux.NewRouter().StrictSlash(true),
// // 		Database: database,
// // 	}
// // }
// // type Controller app.App

// /**
// * Handle the get method to query record and pass the JSON data to front end.
//  */
// func PetGetFunction(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	// param1, ok1 := vars["param1"]
// 	param2, ok2 := vars["param2"]
// 	// if !ok1 {
// 	// 	log.Fatal("No pet in the path")
// 	// }
// 	if !ok2 {
// 		log.Fatal("No ID in the path")
// 	}
// 	// log.Println("param1 is:", param1)
// 	log.Println("param2 is:", param2)

// 	// Query all pets having status, append all pet to pets []Pet array;
// 	pet := &app.Pet
// 	var pets []Pet

// 	// Use the tmpPhotoUrls string to store the &pet.PhotoUrls string templetely, then Decode to []photourl format.
// 	var tmpPhotoUrls string

// 	// Use the tmpTags string to store the &pet.Tags string templetely, then Decode to []tag format.
// 	var tmpTags string

// 	// Using for counting how many records fetched
// 	count := 0
// 	rows, err := app.Database.Query("SELECT id, json_extract(category, '$.id') AS Category_ID, json_extract(category, '$.name') AS Category_Name, name, photoUrls, tags, status FROM `pet` WHERE id = ?", param2)

// 	// Loop the records, do all Types convertion, append the result to pets if no error
// 	for rows.Next() {
// 		// Switch the err and print the outcomes.
// 		switch err := rows.Scan(&pet.ID, &pet.Category.ID, &pet.Category.Name, &pet.Name, &tmpPhotoUrls, &tmpTags, &pet.Status); err {
// 		case sql.ErrNoRows:
// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprintf(w, "Welcome to PetStore!\n")
// 			if err := json.NewEncoder(w).Encode("No rows were returned!"); err != nil {
// 				panic(err)
// 			}
// 			fmt.Fprintf(w, "No rows were returned!\n")
// 		case nil:
// 			// Testing tmpPhotoUrls and tmpTags values getting from the database
// 			fmt.Printf("tmpPhotoUrls from DB is: %s", tmpPhotoUrls)
// 			fmt.Printf("tmpTags from DB is: %s", tmpTags)

// 			// Defining the photoUrlsJSON for Decode the string to []photourl
// 			var photoUrlsJSON []photourl
// 			decodeTmpPhotoUrls := json.NewDecoder(strings.NewReader(tmpPhotoUrls))
// 			errPhotoUrlsJSON := decodeTmpPhotoUrls.Decode(&photoUrlsJSON)

// 			// Testing Decoded values of photoUrlsJSON
// 			fmt.Printf("After Decode tmpPhotoUrls err is %s, array is %s", errPhotoUrlsJSON, photoUrlsJSON)

// 			// Defining the tagsJSON for Decode the string to []photourl
// 			var tagsJSON []tags
// 			decodeTmpTags := json.NewDecoder(strings.NewReader(tmpTags))
// 			errTagsJSON := decodeTmpTags.Decode(&tagsJSON)

// 			// Testing Decoded values of tagsJSON
// 			fmt.Printf("After Decode tmpTags err is %s, array is %s", errTagsJSON, tagsJSON)

// 			if err != nil {
// 				log.Fatal("Database SELECT failed", err)
// 			}

// 			// Append *pet to pets for the final output
// 			pets = append(pets, *pet)
// 			count++
// 		default:
// 			log.Panic("Database SELECT failed: ", err)
// 		}
// 	}
// 	if err != nil {
// 		log.Fatal("Database SELECT failed", err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "There are %v records fetched!\n", count)
// 	log.Println("You fetched pets by id!")
// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(pets); err != nil {
// 		panic(err)
// 	}
// }
