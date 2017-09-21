package mongo

import (
	"strings"
	
	bson "gopkg.in/mgo.v2/bson"

	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

func BsonIdToString(bsonId bson.ObjectId) string {
	tempStr := strings.Replace(bsonId.String(), "ObjectIdHex(", "", -1)
	tempStr = strings.Replace(tempStr, ")", "", -1)
	return tempStr
}

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