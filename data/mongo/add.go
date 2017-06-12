package mongo

import (
	bson "gopkg.in/mgo.v2/bson"
	
	"github.com/quinlanmorake/lib.golang/result"
	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func (mi *MongoInterface) Add(parameters dataGlobals.AddParameters) (result.Result, string) {
	addParameters := parameters.ToBsonMap()
	
	dbSession := mi.MongoSession.Copy()
	defer dbSession.Close()

	collection := dbSession.DB(mi.DbName).C(string(parameters.Table))

	if addError := collection.Insert(addParameters); addError != nil {
		return dataGlobals.GetOperationError(addError.Error()), ""
	} else {
		return result.GetSuccessResult(), addParameters["_id"].(bson.ObjectId).String()
	}
}