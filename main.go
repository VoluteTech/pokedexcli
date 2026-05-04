package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
