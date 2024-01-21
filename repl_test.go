package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		err      error
	}{
		{
			input:    "  ",
			expected: "",
			err:      errors.New("no input was given"),
		},
		{
			input:    "  hello  ",
			expected: "hello",
			err:      nil,
		},
		{
			input:    "  hello  world  ",
			expected: "hello",
			err:      nil,
		},
		{
			input:    "  HellO  World  ",
			expected: "hello",
			err:      nil,
		},
	}

	for _, tc := range cases {
		actual, err := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, tc.expected)
			continue
		}
		if actual != tc.expected {
			t.Errorf("cleanInput(%v) == %v, expected %v", tc.input, actual, tc.expected)
		}
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("error state doesn't match: '%v' vs '%v'", err, tc.err)
			continue
		}
	}
}
