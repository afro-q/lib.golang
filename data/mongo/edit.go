package mongo

import (
	"github.com/quinlanmorake/lib.golang/result"

	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func (mi *MongoInterface) Edit(parameters dataGlobals.EditParameters) result.Result {
	dbSession := mi.MongoSession.Copy()
	defer dbSession.Close()

	collection := dbSession.DB(mi.DbName).C(string(parameters.Table))

	filterParameters := parameters.FilterFields.ToBsonMap()
	fieldsToUpdate := map[string]interface{} {
		"$set": parameters.FieldsToUpdate.ToBsonMap(),
	}
	
	if _, editError := collection.UpdateAll(filterParameters, fieldsToUpdate); editError != nil {
		return dataGlobals.GetOperationError(editError.Error())
	} else {
		return result.GetSuccessResult()
	}
}