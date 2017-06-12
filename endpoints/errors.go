package endpoints

import (
	"github.com/quinlanmorake/lib.golang/result"
)

func InvalidRouteResult(RequestedEndpoint string) result.Result {
	return result.Result {
		ErrorCode: -1,
		ErrorMessage: RequestedEndpoint + " is an invalid endpoint",
	}
}

func InvalidFormPost(errorMsg string) result.Result {
	return result.Result {
		ErrorCode: -2,
		ErrorMessage: "One or more of the form values were invalid. " + errorMsg,
	}
}

func ParametersInvalidOrMissing(errorMsg string) result.Result {
	return result.Result {
		ErrorCode: 1,
		ErrorMessage: "Error with one or more of the required parameters - [" + errorMsg + "].",
	}
}

func RecognitionEncodingError(errorMsg string) result.Result {
	return result.Result {
		ErrorCode: 2,
		ErrorMessage: "Error while encoding the data. " + errorMsg,
	}
}
