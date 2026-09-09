package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cliConfig *config) error {
	// get 20 locations
	URL := ""
	if cliConfig.previousURL != nil {
		URL = *cliConfig.previousURL
	} else {
		fmt.Println("you're on the first page")
		return nil
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
