package main

import (
	"errors"
	"strings"
)

func cleanInput(input string) (string, error) {
	lowerInput := strings.ToLower(input)
	lowerInputs := strings.Fields(lowerInput)
	if len(lowerInputs) == 0 {
		return "", errors.New("no input was given")
	}
	return lowerInputs[0], nil
}
