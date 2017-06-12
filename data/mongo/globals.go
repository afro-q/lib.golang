package mongo

import (
	bson "gopkg.in/mgo.v2/bson"

	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func GetField(mongoDbObjectPropertyName string, mongoDbObjectPropertyValue interface{}) (dbField dataGlobals.Field) {
	dbField = dataGlobals.Field{
		Name: mongoDbObjectPropertyName,
	}
	
	switch mongoDbObjectPropertyValue.(type) {
	case bson.ObjectId:
		dbField.Type = dataGlobals.DbField_Id
		dbField.Value = mongoDbObjectPropertyValue.(bson.ObjectId).Hex()
		break

	case string:
		dbField.Type = dataGlobals.DbField_VarChar
		dbField.Value = mongoDbObjectPropertyValue.(string)
		break
		
	default:
		dbField.Type = dataGlobals.DbField_Unknown
		dbField.Value = mongoDbObjectPropertyValue		
		break
	}

	return
}