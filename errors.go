package goson

import "errors"

var (

	// ErrorIndexOutOfRange represents index out of range error
	ErrorIndexOutOfRange = errors.New("index out of range")

	// ErrorNotArray represents error that target object is not array
	ErrorNotArray = errors.New("not array")

	// ErrorNotMap represents error that target object is not map
	ErrorNotMap = errors.New("not map")

	// ErrorInvalidJSONKey represents error that json key dose not exist
	ErrorInvalidJSONKey = errors.New("invalid json Key")
)
