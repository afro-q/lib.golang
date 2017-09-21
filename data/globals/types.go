package globals

import (
	bson "gopkg.in/mgo.v2/bson"
)

type DatabaseType uint8

const (
	DbType_Mongo = DatabaseType(iota)
)

type Tablename string

type DbField uint8

const (
	DbField_Unknown = DbField(iota)
  DbField_Id 
  DbField_Guid 
  DbField_VarChar
	DbField_Int
	DbField_Blob
	DbField_Object
	DbField_Array
)

type Field struct {
	Name string
	Type DbField
	Value interface{}
}

type DbFieldArray []Field

func (dfa DbFieldArray) ToBsonMap() bson.M {
	mapObject := make(map[string]interface{}, len(dfa))

	for _, field := range dfa {
		if (field.Type == DbField_Id) {
			mapObject[field.Name] = bson.ObjectIdHex(field.Value.(string))
		} else {
			mapObject[field.Name] = field.Value
		} 
	}

	return mapObject
}
