package data

import (
	"github.com/quinlanmorake/lib.golang/result"
	
	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"	
)

func Edit(parameters dataGlobals.EditParameters) result.Result {
	return DatabaseEditImplementors[DatabaseConfig.Type].Edit(parameters)
}
