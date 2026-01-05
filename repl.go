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
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cmd := words[0]
		command, exists := getCommands()[cmd]
		if exists {
			if err := command.callback(); err != nil {
				fmt.Println("Error executing command:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", cmd)
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
