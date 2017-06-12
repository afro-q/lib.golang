package mongo

import (
	"github.com/quinlanmorake/lib.golang/result"

	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func (mi *MongoInterface) Remove(parameters dataGlobals.RemoveParameters) result.Result {
	dbSession := mi.MongoSession.Copy()
	defer dbSession.Close()

	collection := dbSession.DB(mi.DbName).C(string(parameters.Table))
	filterParameters := parameters.FilterFields.ToBsonMap()

	if _, removeError := collection.RemoveAll(filterParameters); removeError != nil {
		return dataGlobals.GetOperationError(removeError.Error())
	} else {
		return result.GetSuccessResult()
	}
}
