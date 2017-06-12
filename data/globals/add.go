package globals

import (
	bson "gopkg.in/mgo.v2/bson"
	
	"github.com/quinlanmorake/lib.golang/result"
)

type AddParameters struct {
	Table Tablename
	Fields []Field	
}

type IAddImplementor interface {
	Setup(config Config) error
	Add(parameters AddParameters) (result.Result, string)
}

func (a AddParameters) ToBsonMap() bson.M {
	mapObject := make(map[string]interface{}, len(a.Fields))

	for _, field := range a.Fields {
		mapObject[field.Name] = field.Value
	}

	// Create the Id
	mapObject["_id"] = bson.NewObjectId()	
	return mapObject
}

