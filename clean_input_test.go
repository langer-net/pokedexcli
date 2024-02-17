package main

import (
	"errors"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
		err      error
	}{
		{
			input:    "  ",
			expected: []string{},
			err:      errors.New("no input was given"),
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
			err:      nil,
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
			err:      nil,
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
			err:      nil,
		},
	}

	for _, tc := range cases {
		actual := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, tc.expected)
			continue
		}
		for i, actualWord := range actual {
			if actualWord != tc.expected[i] {
				t.Errorf("cleanInput(%v) == %v, expected %v", tc.input, actual, tc.expected)
				continue
			}
		}
	}
}
