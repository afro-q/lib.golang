package globals

import (
	"github.com/quinlanmorake/lib.golang/result"
)

func GetOperationError(errorMsg string) result.Result {
	return result.Result{
		ErrorCode: 201,
		ErrorMessage: "An error occured while performing the database operation. " + errorMsg,
	}
}