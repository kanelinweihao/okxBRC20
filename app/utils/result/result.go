package result

import "net/http"

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func GetResultSuccess(data interface{}) (result *Result) {
	result = new(Result)
	result.Code = http.StatusOK
	result.Msg = "OK"
	result.Data = data
	return result
}

func GetResultFail(err error) (result *Result) {
	result = new(Result)
	result.Code = http.StatusBadRequest
	result.Msg = "Fail"
	result.Data = err.Error()
	return result
}
