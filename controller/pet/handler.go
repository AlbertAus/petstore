package pet

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AlbertAus/petstore/model"
)

//Pet defined for Global usuage
type Pet model.Pet

// notFound function handling the Pet Not Found message
func notFound(param2 string, w http.ResponseWriter) {
	// Output Pet Not Found message.
	var errMessage model.PetNotFound
	errMessage.Code, _ = strconv.ParseInt(param2, 10, 64)
	errMessage.Type = "error"
	errMessage.Message = "Pet not found"

	// Write the response text to front end
	outputMessage, outputErr := json.Marshal(errMessage)
	if outputErr == nil {
		http.Error(w, string(outputMessage), 404)
	}
}
