package main

import "github.com/JonasRH355/CLI-Pokedex-in-GO/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	atribute string
}
