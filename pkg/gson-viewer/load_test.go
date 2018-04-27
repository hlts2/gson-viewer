package gsonviewer

import (
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		path           string
		isError        bool
		isGsonInstance bool
	}{
		{
			path:           "../../test_datas/data1.json",
			isError:        false,
			isGsonInstance: true,
		},
		{
			path:           "data1.json",
			isError:        true,
			isGsonInstance: false,
		},
	}

	for i, test := range tests {
		gson, err := LoadJSON(test.path)

		isError := !(err == nil)

		if test.isError != isError {
			t.Errorf("i = %d LoadJSON(path) expected isError: %v, got: %v", i, test.isError, isError)
		}

		isGsonInstance := !(gson == nil)

		if test.isGsonInstance != isGsonInstance {
			t.Errorf("i = %d, LoadJSON(path) expected isGsonInstance: %v, got: %v", i, test.isGsonInstance, isGsonInstance)
		}
	}
}
