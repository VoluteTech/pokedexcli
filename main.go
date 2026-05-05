package main

import (
	"time"

	"github.com/VoluteTech/pokedexcli/internal/api"
)

func main() {
	pokeClient := api.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
