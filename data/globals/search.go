package globals

import (
	bson "gopkg.in/mgo.v2/bson"
	
	"github.com/quinlanmorake/lib.golang/result"	
)

type SearchParameters struct {
	Table Tablename
	SearchFields DbFieldArray
}

type ISearchImplementor interface {
	Setup(config Config) error
	Search(parameters SearchParameters) (result.Result, *Rows)
}

func (s SearchParameters) ToBsonMap() bson.M {
	return s.SearchFields.ToBsonMap()
}

