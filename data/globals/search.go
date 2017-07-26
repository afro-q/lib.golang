package globals

import (
	bson "gopkg.in/mgo.v2/bson"
	
	"github.com/quinlanmorake/lib.golang/result"	
)

type SearchParameters struct {
	Table Tablename
	SearchFields DbFieldArray
	SortFields DbFieldArray
	StartIndex int
	MaxRows int
}

type ISearchImplementor interface {
	Setup(config Config) error
	Search(parameters SearchParameters) (result.Result, *Rows)
}

func (s SearchParameters) ToBsonMap() bson.M {
	return s.SearchFields.ToBsonMap()
}

func (s SearchParameters) GetSortFields() []string {
	sortFields := make([]string, len(s.SortFields))

	for index, _ := range s.SortFields {
		sortFields[index] = s.SortFields[index].Name
	}

	return sortFields
}

