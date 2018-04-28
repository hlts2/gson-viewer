package gsonviewer

import "testing"

func TestNormalizeInputText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "lists.[0].name",
			expected: "lists.0.name",
		},
		{
			input:    "[0]aaaa",
			expected: "0aaaa",
		},
	}

	for i, test := range tests {
		got := NormalizeInputText(&test.input)

		if test.expected != got {
			t.Errorf("i = %d NormalizeInputText(in) expected: %v, got: %v", i, test.expected, got)
		}
	}
}
