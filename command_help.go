package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	availableCommands := getCommands()

	for _, value := range availableCommands {
		fmt.Printf(" - %s: %s\n", value.name, value.description)
	}
	fmt.Println("")
	return nil
}
