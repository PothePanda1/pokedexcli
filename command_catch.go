package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cliConfig *config, args ...string) error {

	pokemon_name := args[0]
	if pokemon_name == "" {
		fmt.Println("Error: Not given a Pokemon")
		return nil
	}
	page, err := cliConfig.pokeapiClient.GetPokemon(pokemon_name)

	if err != nil { // if the url failed and couldn't find a pokemon
		fmt.Println("Error: ", err)
		return err
	} // error isn't really working anyways

	fmt.Println("Throwing a Pokeball at " + pokemon_name + "...")

	number := rand.Intn(page.BaseExp)
	if rand.Intn(number) <= 40 {
		fmt.Println(pokemon_name + " was caught!")
		cliConfig.caughtPokemon[pokemon_name] = page
	} else {
		fmt.Println(pokemon_name + " escaped!")
	}
	return nil
}
