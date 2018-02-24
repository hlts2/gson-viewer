package goson

import "errors"

// Result is search and path result struct
type Result struct {
	object interface{}
}

// Int converts search and path result object to int
func (r *Result) Int() (int, error) {
	v, ok := r.object.(int)
	if !ok {
		return 0, errors.New("")
	}
	return v, nil
}
