package mongo

import (
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	
	"github.com/quinlanmorake/lib.golang/result"
	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func (mi *MongoInterface) Search(parameters dataGlobals.SearchParameters) (result.Result, *dataGlobals.Rows) {
	mongoFilterParameters := parameters.ToBsonMap()
	
	searchSession := mi.MongoSession.Copy()
	defer searchSession.Close()

	collection := searchSession.DB(mi.DbName).C(string(parameters.Table))
	var rowIterator *mgo.Iter

	if len(parameters.SortFields) > 0 {
		fieldsToSortBy := make([]string, len(parameters.SortFields))

		for index, _ := range parameters.SortFields {
			if parameters.SortFields[index].Ascending {
				fieldsToSortBy[index] = parameters.SortFields[index].Name
			} else {
				fieldsToSortBy[index] = "-" + parameters.SortFields[index].Name
			}
		}
		
		rowIterator = collection.Find(mongoFilterParameters).Sort(fieldsToSortBy...).Skip(parameters.StartIndex).Limit(parameters.MaxRows).Iter()
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