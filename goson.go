package goson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (

	// ErrorIndexOutOfRange represents index out of range error
	ErrorIndexOutOfRange = errors.New("index out of range")

	// ErrorNotArray represents error that target object is not array
	ErrorNotArray = errors.New("not array")

	// ErrorNotMap represents error that target object is not map
	ErrorNotMap = errors.New("not map")

	// ErrorInvalidJSONKey represents error that json key dose not exist
	ErrorInvalidJSONKey = errors.New("invalid json Key")

	// ErrorInvalidSyntax represents invaild syntax error
	ErrorInvalidSyntax = errors.New("invalid syntax")
)

// GosonError represents a json parse error
type GosonError struct {
	Value interface{}
	Err   error
}

func (e *GosonError) Error() string {
	return Quote(e.Value) + ": " + e.Err.Error()
}

// ResultError represents a conversion error
type ResultError struct {
	Fn     string
	Object interface{}
	Err    error
}

func (e *ResultError) Error() string {
	return "goson." + e.Fn + ": parsing " + Quote(e.Object) + ": " + e.Err.Error()
}

// Quote returns quoted object string
func Quote(object interface{}) string {
	return fmt.Sprintf("\"%v\"", object)
}

// Result represents a json value that is returned from Search() and Path().
type Result struct {
	object interface{}
}

// Goson is goson base struct
type Goson struct {
	jsonObject interface{}
}

// NewGosonFromByte returns Goson instance created from byte array
func NewGosonFromByte(data []byte) (*Goson, error) {
	g := new(Goson)

	if err := decode(bytes.NewReader(data), &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

// NewGosonFromString returns Goson instance created from string
func NewGosonFromString(data string) (*Goson, error) {
	g := new(Goson)

	if err := decode(strings.NewReader(data), &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

// NewGosonFromReader returns Goson instance created from io.Reader
func NewGosonFromReader(reader io.Reader) (*Goson, error) {
	g := new(Goson)

	if err := decode(reader, &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

func decode(reader io.Reader, object *interface{}) error {
	dec := json.NewDecoder(reader)
	if err := dec.Decode(object); err != nil {
		return err
	}
	return nil
}

// Indent converts json object to json string
func (g *Goson) Indent(prefix, indent string) (string, error) {
	return indentJSONString(g.jsonObject, prefix, indent)
}

func indentJSONString(object interface{}, prefix, indent string) (string, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, data, prefix, indent); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Search returns json value corresponding to keys. keys represents key of hierarchy of json
func (g *Goson) Search(keys ...string) (*Result, error) {
	var err error

	jsonObject := g.jsonObject

	for _, key := range keys {
		if jsonObject, err = search(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return &Result{jsonObject}, nil
}

// Path returns json value corresponding to path.
func (g *Goson) Path(path string) (*Result, error) {
	var err error

	jsonObject := g.jsonObject

	for _, key := range strings.Split(path, ".") {
		if jsonObject, err = search(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return &Result{jsonObject}, nil
}

func search(object interface{}, key string) (interface{}, error) {
	const fn = "search"

	index, err := strconv.Atoi(key)
	if err == nil {
		switch object.(type) {
		case []interface{}:
		default:
			return nil, &GosonError{object, ErrorNotArray}
		}

		v := object.([]interface{})

		if 0 <= index && index < len(v) {
			return v[index], nil
		}

		return nil, &GosonError{index, ErrorIndexOutOfRange}
	}

	switch object.(type) {
	case map[string]interface{}:
	default:
		return nil, &GosonError{object, ErrorNotArray}
	}

	v, ok := object.(map[string]interface{})[key]
	if !ok {
		return nil, &GosonError{key, ErrorInvalidJSONKey}
	}

	return v, nil
}

// Uint8 converts an interface{} to a uint8 and returns an error if types don't match.
func (r *Result) Uint8() (uint8, error) {
	const fn = "uint8"

	switch r.object.(type) {
	case uint8:
		return r.object.(uint8), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint16 converts an interface{} to a uint16 and returns an error if types don't match.
func (r *Result) Uint16() (uint16, error) {
	const fn = "uint16"

	switch r.object.(type) {
	case uint16:
		return r.object.(uint16), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint32 converts an interface{} to a uint32 and returns an error if types don't match.
func (r *Result) Uint32() (uint32, error) {
	const fn = "uint32"

	switch r.object.(type) {
	case uint32:
		return r.object.(uint32), nil
	default:
		return 0, &ResultError{fn, r.object, ErrorInvalidSyntax}
	}
}

// Uint64 converts an interface{} to a uint64 and returns an error if types don't match.
func (r *Result) Uint64() (uint64, error) {
	const fn = "uint64"

	switch r.object.(type) {
	case uint64:
		return r.object.(uint64), nil
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
