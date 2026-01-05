package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello     world   ",
			expected: []string{"hello", "world"},
		}, {
			input:    "GoLang is FUN",
			expected: []string{"golang", "is", "fun"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expextedWord := c.expected[i]
			if word != expextedWord {
				t.Errorf("cleanInput(%q) = %q; want %q", c.input, actual, c.expected)
			}
		}
	}
}
