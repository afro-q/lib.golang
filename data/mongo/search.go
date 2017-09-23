package mongo

import (
	//mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	
	"github.com/quinlanmorake/lib.golang/result"
	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func (mi *MongoInterface) Search(parameters dataGlobals.SearchParameters) (result.Result, *dataGlobals.Rows) {
	mongoFilterParameters := parameters.ToBsonMap()
	
	searchSession := mi.MongoSession.Copy()
	defer searchSession.Close()

	collection := searchSession.DB(mi.DbName).C(string(parameters.Table))

	var mongoQuery = collection.Find(mongoFilterParameters)	
	if len(parameters.SortFields) > 0 {
		fieldsToSortBy := make([]string, len(parameters.SortFields))

		for index, _ := range parameters.SortFields {
			if parameters.SortFields[index].Ascending {
				fieldsToSortBy[index] = parameters.SortFields[index].Name
			} else {
				fieldsToSortBy[index] = "-" + parameters.SortFields[index].Name
			}
		}

		mongoQuery = mongoQuery.Sort(fieldsToSortBy...)
	} 

	if parameters.StartIndex > 0 {
		mongoQuery = mongoQuery.Skip(parameters.StartIndex)
	}

	if parameters.MaxRows > 0 {
		mongoQuery = mongoQuery.Limit(parameters.MaxRows)
	}

	rowIterator := mongoQuery.Iter()
	
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