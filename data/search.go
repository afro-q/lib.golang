package data

import (
	"github.com/quinlanmorake/lib.golang/result"

	dataGlobals "github.com/quinlanmorake/lib.golang/data/globals"	
)

func Search(parameters dataGlobals.SearchParameters) (result.Result, *dataGlobals.Rows) {
	return DatabaseSearchImplementors[DatabaseConfig.Type].Search(parameters)
}
