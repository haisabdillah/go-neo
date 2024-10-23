package errors

type Err struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors,omitempty"`
}

func (e Err) Error() string {
	return e.Message
}

func InvalidParam() Err {
	return Err{400, "Invalid param", nil}
}

func InvalidJson(err error) Err {
	return Err{400, "Invalid bind json", nil}
}

func BadRequest(data interface{}) Err {
	return Err{400, "Bad Request", data}
}

func Validation(data interface{}) Err {
	return Err{400, "Invalid validation", data}
}

func NotFound(str string) Err {
	return Err{404, str, nil}
}

func Unauthenticate(str string) Err {
	return Err{401, str, nil}
}

func InternalServer(err error) Err {
	return Err{500, "Internal server error", err}
}
