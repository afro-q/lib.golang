package mongo

import (
	mgo "gopkg.in/mgo.v2"

	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"
)

type MongoInterface struct {
	MongoSession *mgo.Session
	DbName string
}

func (mi *MongoInterface) Setup(config dataGlobals.Config) error {
	if session, initError := mgo.Dial(config.Host); initError != nil {
		return initError
	} else {
		mi.MongoSession = session
		mi.DbName = config.Database
	}
	
	return mi.MongoSession.Ping()
}
