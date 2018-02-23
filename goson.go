package goson

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

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
