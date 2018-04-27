package gsonviewer

import "testing"

func TestNormalizeInputText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "list.[0].name",
			expected: "list.0.name",
		},
	}

	for i, test := range tests {
		got := NormalizeInputText(&test.input)

		if test.expected != got {
			t.Errorf("i = %d NormalizeInputText(in) expected: %v, got: %v", i, test.expected, got)
		}
	}
}
