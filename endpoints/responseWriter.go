package endpoints

import (
	"encoding/json"
	"io"
)

func WriteResponse(writer io.Writer, responseData interface{}) {
	jsonWriter := json.NewEncoder(writer)
	jsonWriter.Encode(responseData)
}