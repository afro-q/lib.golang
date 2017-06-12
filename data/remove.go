package data

import (
	"github.com/quinlanmorake/lib.golang/result"

	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"	
)

func Remove(parameters dataGlobals.RemoveParameters) result.Result {
	return DatabaseRemoveImplementors[DatabaseConfig.Type].Remove(parameters)
}