package mongo

import (
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	
	"github.com/quinlanmorake/lib.golang/result"
	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func (mi *MongoInterface) Search(parameters dataGlobals.SearchParameters) (result.Result, *dataGlobals.Rows) {
	mongoFilterParameters := parameters.ToBsonMap()
	sortParameters := parameters.GetSortFields()
	
	searchSession := mi.MongoSession.Copy()
	defer searchSession.Close()

	collection := searchSession.DB(mi.DbName).C(string(parameters.Table))
	var rowIterator *mgo.Iter

	if len(sortParameters) > 0 {
		rowIterator = collection.Find(mongoFilterParameters).Sort(sortParameters...).Skip(parameters.StartIndex).Limit(parameters.MaxRows).Iter()
	} else {
		rowIterator = collection.Find(mongoFilterParameters).Iter()
	}
	
	var queryResultEntry interface{}
	rowsFound := dataGlobals.Rows{}
	for rowIterator.Next(&queryResultEntry) {
		currentRow := dataGlobals.Row{}

		for key, value := range queryResultEntry.(bson.M) {
			currentRow = append(currentRow, GetField(key, value))
		}		
		rowsFound = append(rowsFound, currentRow)
	}

	if findError :=	rowIterator.Close(); findError != nil {
		return dataGlobals.GetOperationError(findError.Error()), nil
	}

	return result.GetSuccessResult(), &rowsFound
}