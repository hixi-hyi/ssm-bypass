package main

import "testing"

func TestEscape(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "foo", expected: "foo"},
		{input: "foo\nbar", expected: "foo\\nbar"},
		{input: "foo\"bar", expected: "foo\\\"bar"},
		{input: "foo'bar", expected: "foo'bar"},
		{input: "foo$bar", expected: "foo$bar"},
		{input: "foo`bar", expected: "foo`bar"},
	}

	for _, test := range tests {
		actual := escape(test.input)
		if actual != test.expected {
			t.Errorf("escape(%s): expected %s, actual %s", test.input, test.expected, actual)
		}
	}
}
