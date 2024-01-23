package main

import (
	"github.com/langer-net/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.CreateNewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
