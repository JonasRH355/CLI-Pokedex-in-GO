package main

type cliCommand struct {
	name string
	description string
	callback func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next location areas's name",
			callback:    commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display the previous location areas's name",
			callback: commandMapB,
		},
	}
}