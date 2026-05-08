package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Prev

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Prev

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func explore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func catch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonInfos(name)
	if err != nil {
		return err
	}
	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s has escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s has been captured\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}

func inspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	if _, ok := cfg.caughtPokemon[name]; !ok {
		return errors.New("you didn't catch this pokemon")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemonInfos(name)
	if err != nil {
		return err
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("- %v\n", t.Type.Name)
	}

	return nil
}

