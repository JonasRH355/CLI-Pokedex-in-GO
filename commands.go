package main

import (
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandMap(cfg *config) error {
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

func commandMapB(cfg *config) error {
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

func commandExplore(cfg *config) error {
	val, err := cfg.pokeapiClient.GetExploreRegion(cfg.atribute)
	if err != nil {
		return err
	}
	
	for i := 0; i < len(val.Pokemon_encounters); i++ {
		fmt.Println(val.Pokemon_encounters[i].Pokemon.Name)
	}

	return nil
}