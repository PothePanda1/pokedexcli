package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location struct {
	Name string `json:"name"`
}

type page struct {
	Results  []location `json:"results"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
}

func commandMap(cliConfig *config) error {
	// get 20 locations

	URL := "https://pokeapi.co/api/v2/location-area/"

	if cliConfig.nextURL != nil {
		URL = *cliConfig.nextURL
	}

	res, err := http.Get(URL)
	if err != nil { // if the url failed and couldn't find a location area
		fmt.Println("Error: ", err)
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body) // convert the response type to []byte
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	current := page{}
	if err := json.Unmarshal(data, &current); err != nil {
		return err
	}
	cliConfig.nextURL = current.Next
	cliConfig.previousURL = current.Previous
	for _, location := range current.Results {
		fmt.Println(location.Name)
	}
	return nil
}
