package main

import (
	"github.com/langer-net/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.CreateNewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
