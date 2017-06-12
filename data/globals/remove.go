package globals

import (
	"github.com/quinlanmorake/lib.golang/result"
)

type RemoveParameters struct {
	Table Tablename
	FilterFields DbFieldArray
}

type IRemoveImplementor interface {
	Setup(config Config) error
	Remove(parameters RemoveParameters) result.Result
}
