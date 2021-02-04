package utils

type MyError struct {
	Code    int
	Message string
}

func (m MyError) Error() string {
	return m.Message
}

func NewMyError(code int, msg string) error {
	return &MyError{
		Code:    code,
		Message: msg,
	}
}
