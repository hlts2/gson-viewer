package goson

import (
	"errors"
	"fmt"
)

//ErrorInvalidSyntax represents invaild syntax error
var ErrorInvalidSyntax = errors.New("invalid syntax")

//ResultError represents a conversion error
type ResultError struct {
	Fn     string
	Object interface{}
	Err    error
}

func (e *ResultError) Error() string {
	return "goson." + e.Fn + ": parsing: " + Quote(e.Object) + ": " + e.Err.Error()
}

//Quote returns quoted value
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
