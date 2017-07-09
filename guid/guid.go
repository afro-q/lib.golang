package guid

import (
	"fmt"
	
	"github.com/satori/go.uuid"
)

func NewGuid() string {
	return fmt.Sprintf("%v", uuid.NewV4())	
}