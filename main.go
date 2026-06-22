package main

import (
	"time"
	"github.com/JonasRH355/CLI-Pokedex-in-GO/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}

