package endpoints

import (
	"encoding/json"	
	"net/http"

	"github.com/quinlanmorake/lib.golang/result"
)

// This will return a function that checks if the POST form is valid, and calls the required handler if so
func ValidateAndHandle(handler http.HandlerFunc) http.HandlerFunc {
	return func (responseWriter http.ResponseWriter, request *http.Request) {
		if formParseError := request.ParseForm(); formParseError != nil {
			result := InvalidFormPost(formParseError.Error())
			WriteResponse(responseWriter, result)
			return
		}
	
		handler(responseWriter, request)
	}
}

// This will return a function that parses a multi-part form, checking validatity, and calling the required handle if valid
func ValidateMultipartAndHandle(handler http.HandlerFunc) http.HandlerFunc {
	return func (responseWriter http.ResponseWriter, request *http.Request) {
		if formParseError := request.ParseMultipartForm(32 << 30); formParseError != nil {
			result := InvalidFormPost(formParseError.Error())
			WriteResponse(responseWriter, result)
			return
		}

		handler(responseWriter, request)
	}
}

// This will parse a JSON body into its corresponding object
func ParseJsonBody(request *http.Request, target interface{}) result.Result {
	defer request.Body.Close()
	
	decoder := json.NewDecoder(request.Body)
	if parseError := decoder.Decode(target); parseError != nil {
		return InvalidFormPost(parseError.Error())
	} else {
		return result.GetSuccessResult()
	}
}