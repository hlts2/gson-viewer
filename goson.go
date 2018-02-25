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
)

// GosonError represents a json parse error
type GosonError struct {
	Fn    string
	Value interface{}
	Err   error
}

func (e *GosonError) Error() string {
	return "goson." + e.Fn + ": parsing " + Quote(e.Value) + ": " + e.Err.Error()
}

// Quote returns quoted object string
func Quote(object interface{}) string {
	return fmt.Sprintf("\"%v\"", object)
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
			return nil, &GosonError{fn, object, ErrorNotArray}
		}

		v := object.([]interface{})

		if 0 <= index && index < len(v) {
			return v[index], nil
		}

		return nil, ErrorIndexOutOfRange
	}

	switch object.(type) {
	case map[string]interface{}:
	default:
		return nil, &GosonError{fn, object, ErrorNotArray}
	}

	v, ok := object.(map[string]interface{})[key]
	if !ok {
		return nil, ErrorInvalidJSONKey
	}

	return v, nil
}
