package result

type Result struct {
	ErrorCode float64 `json:"code"`
	ErrorMessage string `json:"message"`
}

func (r Result) IsOk() bool {
	return 0 == r.ErrorCode
}

func (r Result) IsNotOk() bool {
	return 0 != r.ErrorCode
}

func GetSuccessResult() Result {
	return Result {
		ErrorCode: 0,
		ErrorMessage: "",
	}
}