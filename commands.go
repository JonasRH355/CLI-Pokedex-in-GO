package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, arg ...string) error {
	fmt.Println("------------------------------")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("------------------------------")
	return nil
}

func commandExit(cfg *config, arg ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandMap(cfg *config, arg ...string) error {
	val, err := cfg.pokeapiClient.GetMap(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	
	for i := 0; i < len(val.Results); i++ {
		fmt.Println(val.Results[i].Name)
	}

	cfg.nextLocationsURL = val.Next
	cfg.prevLocationsURL = val.Previous
	
	return nil
}

func commandMapB(cfg *config, arg ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	val, err := cfg.pokeapiClient.GetMap(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	
	for i := 0; i < len(val.Results); i++ {
		fmt.Println(val.Results[i].Name)
	}

	
	cfg.nextLocationsURL = val.Next
	cfg.prevLocationsURL = val.Previous
	
	return nil
}

func commandExplore(cfg *config, arg ...string) error {
	if len(arg) != 1 {
		return fmt.Errorf("you must provide a location name")
	}

	val, err := cfg.pokeapiClient.GetExploreRegion(arg[0])
	if err != nil {
		return err
	}
	
	fmt.Printf("Exploring %s...\n", val.Name)
	fmt.Println("Found Pokemon: ")
	for i := 0; i < len(val.Pokemon_encounters); i++ {
		fmt.Println(val.Pokemon_encounters[i].Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, arg ...string) error {
	if len(arg) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", arg[0])

	val, err := cfg.pokeapiClient.GetPokemon(arg[0])
	if err != nil {
		return err
	}
	
	
	if cfg.lastCommand != arg[0] {
		cfg.lastCommand = arg[0]
		cfg.repeat = 1
	} else {
		cfg.repeat += 1
	}

	catche := rand.Intn(val.Base_experience) 
	if catche < 10*cfg.repeat {
		fmt.Printf("%s was caught!\n", arg[0])
		cfg.repeat = 0
		cfg.pokedex[fmt.Sprintf("%s_%v",val.Name,len(cfg.pokedex))] = val
	} else {
		fmt.Printf("%s escaped!\n", arg[0])
	}
	
	return nil
}

func commandPokedex(cfg *config, arg ...string) error {
	for pokemons := range cfg.pokedex {
		fmt.Println(pokemons)
	}
	
	return nil
}

func commandInspect(cfg *config, arg ...string) error {
	if len(arg) != 1 {
		return fmt.Errorf("you must provide a pokemon nick name")
	}
	
	name := arg[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}