package globals

import (
	"encoding/json"
	
	"github.com/quinlanmorake/lib.golang/result"
)

type Row []Field

type Rows []Row

func FieldSerializeError(error string) result.Result {
	return result.Result {
		ErrorCode: 200,
		ErrorMessage: "Error converting data from database to the required object. " + error,
	}
}

func (r Row) Parse(target interface{}) (parseResult result.Result) {
	parseResult = result.GetSuccessResult()
	
	targetMap := make(map[string]interface{}, 0)
	for _, entry := range r {
		targetMap[entry.Name] = entry.Value
	}

	if byteArray, convertToJsonError := json.Marshal(targetMap); convertToJsonError != nil {
		parseResult = FieldSerializeError(convertToJsonError.Error())
	} else {
		if convertFromJsonError := json.Unmarshal(byteArray, target); convertFromJsonError != nil {
			parseResult = FieldSerializeError(convertFromJsonError.Error())
		}
	}

	return
}

//func (r Rows) Parse(target interface{}) { }