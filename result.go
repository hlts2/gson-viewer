package goson

import (
	"errors"
	"fmt"
)

// ErrorInvalidSyntax represents invaild syntax error
var ErrorInvalidSyntax = errors.New("invalid syntax")

// ResultError represents a conversion error
type ResultError struct {
	Fn     string
	Object interface{}
	Err    error
}

func (e *ResultError) Error() string {
	return "goson." + e.Fn + ": parsing: " + Quote(e.Object) + ": " + e.Err.Error()
}

// Quote returns quoted object string
func Quote(object interface{}) string {
	return fmt.Sprintf("\"%v\"", object)
}

// Result represents a json value that is returned from Search() and Path().
type Result struct {
	object interface{}
}

// Int converts an interface{} to a int and returns an error if types don't match.
func (r *Result) Int() (int, error) {
	const fn = "Int"

	switch r.object.(type) {
	case int:
		return r.object.(int), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int8 converts an interface{} to a int8 and returns an error if types don't match.
func (r *Result) Int8() (int8, error) {
	const fn = "int8"

	switch r.object.(type) {
	case int8:
		return r.object.(int8), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int16 converts an interface{} to a int16 and returns an error if types don't match.
func (r *Result) Int16() (int16, error) {
	const fn = "int16"

	switch r.object.(type) {
	case int16:
		return r.object.(int16), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int32 converts an interface{} to a int32 and returns an error if types don't match.
func (r *Result) Int32() (int32, error) {
	const fn = "Int32"

	switch r.object.(type) {
	case int32:
		return r.object.(int32), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Int64 converts an interface{} to a int64 and returns an error if types don't match.
func (r *Result) Int64() (int64, error) {
	const fn = "Int64"

	switch r.object.(type) {
	case int64:
		return r.object.(int64), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// float32 converts an interface{} to a float32 and returns an error if types don't match.
func (r *Result) float32() (float32, error) {
	const fn = "float32"

	switch r.object.(type) {
	case float32:
		return r.object.(float32), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// float64 converts an interface{} to a float64 and returns an error if types don't match.
func (r *Result) float64() (float64, error) {
	const fn = "float64"

	switch r.object.(type) {
	case float64:
		return r.object.(float64), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}
