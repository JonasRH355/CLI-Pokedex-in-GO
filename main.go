package main

import (
	"time"
	"github.com/JonasRH355/CLI-Pokedex-in-GO/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		lastCommand: "",
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}

