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
	if _, err := NewGosonFromString(testData1); err != nil {
		t.Errorf("NewGosonFromString is err: %v", err)
	}
}
