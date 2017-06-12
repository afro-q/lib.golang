package globals

import (
	"github.com/quinlanmorake/lib.golang/result"
)

type EditParameters struct {
	Table Tablename
	
	FilterFields DbFieldArray
	FieldsToUpdate DbFieldArray
}

type IEditImplementor interface {
	Setup(config Config) error
	
	Edit(parameters EditParameters) result.Result
}
