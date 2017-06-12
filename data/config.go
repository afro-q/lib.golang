package data

import (
	"github.com/quinlanmorake/lib.golang/data/mongo"
	
	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"	
)

var DatabaseConfig dataGlobals.Config

var DatabaseSearchImplementors map[dataGlobals.DatabaseType]dataGlobals.ISearchImplementor
var DatabaseAddImplementors map[dataGlobals.DatabaseType]dataGlobals.IAddImplementor
var DatabaseEditImplementors map[dataGlobals.DatabaseType]dataGlobals.IEditImplementor
var DatabaseRemoveImplementors map[dataGlobals.DatabaseType]dataGlobals.IRemoveImplementor

var mongoInterface mongo.MongoInterface

func SetConfig(config dataGlobals.Config) (setupError error) {
	DatabaseConfig = config
	mongoInterface = mongo.MongoInterface{}

	if setupError = mongoInterface.Setup(DatabaseConfig); setupError != nil {
		return
	}
	
	if setupError = setupSearchers(); setupError != nil {
		return
	}

	if setupError = setupAdders(); setupError != nil {
		return
	}

	if setupError = setupEditers(); setupError != nil {
		return
	}

	if setupError = setupRemovers(); setupError != nil {
		return
	}
	
	return
}

func setupSearchers() error {
	DatabaseSearchImplementors = make(map[dataGlobals.DatabaseType]dataGlobals.ISearchImplementor, 1)
	DatabaseSearchImplementors[dataGlobals.DbType_Mongo] = &mongoInterface
	
	return nil
}

func setupAdders() error {
	DatabaseAddImplementors = make(map[dataGlobals.DatabaseType]dataGlobals.IAddImplementor, 1)
	DatabaseAddImplementors[dataGlobals.DbType_Mongo] = &mongoInterface
	
	return nil
}

func setupEditers() error {
	DatabaseEditImplementors = make(map[dataGlobals.DatabaseType]dataGlobals.IEditImplementor, 1)
	DatabaseEditImplementors[dataGlobals.DbType_Mongo] = &mongoInterface
	
	return nil
}

func setupRemovers() error {
	DatabaseRemoveImplementors = make(map[dataGlobals.DatabaseType]dataGlobals.IRemoveImplementor, 1)
	DatabaseRemoveImplementors[dataGlobals.DbType_Mongo] = &mongoInterface
	
	return nil
}
