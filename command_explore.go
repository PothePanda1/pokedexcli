package main

import (
	"fmt"
)

func commandExplore(cliConfig *config, args ...string) error {

	location_name := args[0]
	page, err := cliConfig.pokeapiClient.GetLocation(location_name)

	if err != nil { // if the url failed and couldn't find a location area
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Exploring " + location_name + "...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range page.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
