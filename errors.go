package goson

import "errors"

var (

	// ErrorIndexOutOfRange is index out of range error
	ErrorIndexOutOfRange = errors.New("index out of range")

	// ErrorNotArray is error that target object is not array
	ErrorNotArray = errors.New("not array")

	// ErrorNotMap is error that target object is not map
	ErrorNotMap = errors.New("not map")

	// ErrorInvalidJSONKey is error that json path dose not exist
	ErrorInvalidJSONKey = errors.New("invalid json path")
)
