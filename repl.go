package main

import "strings"

func cleanInput(text string) []string {
	parts := strings.Fields(text)
	for i := range parts {
		lowered := strings.ToLower(parts[i])
		parts[i] = lowered
	}

	return parts
}
