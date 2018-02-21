package goson

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

var (

	// ErrIndexOufOfBounds is index out of bounds
	ErrIndexOufOfBounds = errors.New("")
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
	return nil, nil
}

func search(object interface{}, key string) (interface{}, error) {
	index, err := strconv.Atoi(key)
	if err == nil {
		switch object.(type) {
		case []interface{}:
		default:
			return nil, errors.New("")
		}

		if len(object.([]interface{})) > index && index > 0 {
			return object.([]interface{})[index], nil
		}

		return nil, ErrIndexOufOfBounds
	}

	switch object.(type) {
	case map[string]interface{}:
	default:
		return nil, ErrIndexOufOfBounds
	}

	v, ok := object.(map[string]interface{})
	if !ok {
		return nil, errors.New("")
	}

	return v, nil
}
