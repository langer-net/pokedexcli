package main

import (
	"strings"
)

func cleanInput(input string) []string {
	lowerInput := strings.ToLower(input)
	words := strings.Fields(lowerInput)
	return words
}
