package main

type cliCommand struct {
	name string
	description string
	callback func(cfg *config, arg ...string) error
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
		"explore": {
			name: "explore",
			description: "Show every pokemon how the player could catch at that place",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Try to catch a pokemon",
			callback: commandCatch,
		},
		"pokedex": {
			name: "pokedex",
			description: "Show all the pokemons you already have",
			callback: commandPokedex,
		},
	}
}