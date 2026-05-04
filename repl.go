package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			clearedInput := cleanInput(userInput)
			fmt.Printf("Your command was: %v\n", clearedInput[0])
		}
	}
}

func cleanInput(text string) []string {
	parts := strings.Fields(text)
	for i := range parts {
		lowered := strings.ToLower(parts[i])
		parts[i] = lowered
	}

	return parts
}
