package goson

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

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

// Goson is goson base struct
type Goson struct {
	jsonObject interface{}
}

// NewGoson returns Goson instance
func NewGoson(data []byte) (*Goson, error) {
	g := new(Goson)

	if err := json.Unmarshal(data, &g.jsonObject); err != nil {
		return nil, err
	}
	return g, nil
}

// StringIndent converts json object to pretty json string
func (g *Goson) StringIndent(prefix, indent string) (string, error) {
	jsonData, err := json.Marshal(g.jsonObject)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, jsonData, prefix, indent); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Search returns json value corresponding to keys. keys represents key of hierarchy of json
func (g *Goson) Search(keys ...string) (interface{}, error) {
	var err error
	jsonObject := g.jsonObject

	for _, key := range keys {
		if jsonObject, err = search(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return jsonObject, nil
}

// Path returns json value corresponding to path.
func (g *Goson) Path(path string) (interface{}, error) {
	var err error
	jsonObject := g.jsonObject

	for _, key := range strings.Split(path, ".") {
		if jsonObject, err = search(jsonObject, key); err != nil {
			return nil, err
		}
	}
	return jsonObject, nil
}

func search(object interface{}, key string) (interface{}, error) {
	index, err := strconv.Atoi(key)
	if err == nil {
		switch object.(type) {
		case []interface{}:
		default:
			return nil, ErrorNotArray
		}

		if len(object.([]interface{})) > index && index > 0 {
			return object.([]interface{})[index], nil
		}

		return nil, ErrorIndexOutOfRange
	}

	switch object.(type) {
	case map[string]interface{}:
	default:
		return nil, ErrorNotMap
	}

	mmap, _ := object.(map[string]interface{})

	v, ok := mmap[key]
	if !ok {
		return nil, ErrorInvalidJSONKey
	}

	return v, nil
}
