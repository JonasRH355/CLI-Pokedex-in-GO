package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

		
)



func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		
		commandToken := input[0]

		if len(input) > 1 {
			cfg.atribute = input[1]
		}
		 
		command, exist := getCommands()[commandToken]
		if exist {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

