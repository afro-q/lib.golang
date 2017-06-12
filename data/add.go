package data

import (
	"github.com/quinlanmorake/lib.golang/result"
	"github.com/quinlanmorake/lib.golang/data/globals"	
)

func Add(parameters globals.AddParameters) (result.Result, string) {
	return DatabaseAddImplementors[DatabaseConfig.Type].Add(parameters)
}
