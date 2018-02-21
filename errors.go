package goson

// ErrorIndexOutOfRange is index out of range error
type ErrorIndexOutOfRange struct {
}

func (err ErrorIndexOutOfRange) Error() string {
	return "index out of range"
}

// ErrorInvalidJSONPath is invalid json path error
type ErrorInvalidJSONPath struct {
}

func (err ErrorInvalidJSONPath) Error() string {
	return "invalid json path"
}
