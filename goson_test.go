package goson

import (
	"testing"
)

var testData1 = `
{}
`

var testData2 = `
{
    "foo": 1,
    "baz": 123.1,
    "array": [
        {
            "foo": 1
        },
        {
            "bar": 2
        }
    ],
    "sub": {
        "foo": 1,
        "array": [
            1,
            2,
            3
        ]
    },
    "bool": true
}
`

var testData3 = `
[{"name": "hlts2"}, {"name": "hiroto"}]
`
var testData4 = `
hello world
`

func TestNew(t *testing.T) {
	if _, err := NewGoson([]byte(testData1)); err != nil {
		t.Errorf("NewGoson is error: %v", err)
	}

	if _, err := NewGoson([]byte(testData2)); err != nil {
		t.Errorf("NewGoson is error: %v", err)
	}

	if _, err := NewGoson([]byte(testData3)); err != nil {
		t.Errorf("NewGoson is error: %v", err)
	}

	if _, err := NewGoson([]byte(testData4)); err == nil {
		t.Errorf("NewGoson is not error: %v", err)
	}
}

func TestSearch(t *testing.T) {
	g, _ := NewGoson([]byte(testData2))
	if result, err := g.Search("foo"); err == nil {
		if result != 1 {
			t.Errorf("Expecting 1, got %v", result)
		}
	} else {
		t.Errorf("Search is error: %v", err)
	}
}
