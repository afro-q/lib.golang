package guid

import (
	"fmt"
	
	"github.com/satori/go.uuid"
)

func NewGuid() string {
	var newGuid string
	fmt.Sprintf(newGuid, uuid.NewV4())
	
	return newGuid
}