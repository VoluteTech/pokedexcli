package main

import (
	"time"

	"github.com/VoluteTech/pokedexcli/internal/api"
)

func main() {
	pokeClient := api.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
