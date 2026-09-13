package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type encounter struct {
	Pokemon resourceName `json:"pokemon"` // grabs the name from the pokemon list
}

type e_page struct {
	PokemonEncounters []encounter `json:"pokemon_encounters"`
}

// Similar to location_list.go but retrieves specific area instead of starting from id 0
func (c *Client) GetLocation(locationName string) (e_page, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		page := e_page{}
		err := json.Unmarshal(val, &page)
		if err != nil {
			return e_page{}, err
		}

		return page, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil { // if the url failed and couldn't find a location area
		return e_page{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return e_page{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil { // convert the response type to []byte
		return e_page{}, err
	}

	pg := e_page{}
	err = json.Unmarshal(dat, &pg)
	if err != nil {
		return e_page{}, err
	}

	return pg, nil
}
