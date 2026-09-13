package main

import (
	"errors"
	"fmt"
)

func commandMap(cliConfig *config, _ ...string) error {

	page, err := cliConfig.pokeapiClient.ListLocations(cliConfig.nextURL)

	if err != nil { // if the url failed and couldn't find a location area
		fmt.Println("Error: ", err)
		return err
	}
	cliConfig.nextURL = page.Next
	cliConfig.previousURL = page.Previous
	for _, location := range page.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cliConfig *config, _ ...string) error {
	if cliConfig.previousURL == nil {
		return errors.New("you're on the first page")
	}

	// exact same thing as Map, but just going backwards
	page, err := cliConfig.pokeapiClient.ListLocations(cliConfig.previousURL)

	if err != nil { // if the url failed and couldn't find a location area
		fmt.Println("Error: ", err)
		return err
	}
	cliConfig.nextURL = page.Next
	cliConfig.previousURL = page.Previous
	for _, location := range page.Results {
		fmt.Println(location.Name)
	}
	return nil
}
