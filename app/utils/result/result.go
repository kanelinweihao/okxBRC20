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
