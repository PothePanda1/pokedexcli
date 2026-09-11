package main

import (
	"time"

	"github.com/PoThePanda1/pokedexcli/internal/pokeapi"
)

func main() {
	pokeclient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cliConfig := &config{
		registry:      getCommands(),
		pokeapiClient: pokeclient,
	}
	startRepl(cliConfig)
}
