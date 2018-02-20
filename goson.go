package goson

import (
	"bytes"
	"encoding/json"
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
	/*
		var object interface{}

			for _, key := range keys {
				if mmap, ok := g.jsonObject.(map[string]interface{}); ok {

				} else if marray, ok := g.jsonObject.([]interface{}); ok {

				} else {
					return nil, errors.New("")
				}
			}
	*/

	return nil, nil
}
