package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type StatDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonStat struct {
	BaseStat int        `json:"base_stat"`
	Effort   int        `json:"effort"`
	Stat     StatDetail `json:"stat"`
}

type Pokemon struct { // page is more for batch results tbh
	Name    string        `json:"name"`
	BaseExp int           `json:"base_experience"`
	Weight  int           `json:"weight"`
	Height  int           `json:"height"`
	Stats   []PokemonStat `json:"stats"`
	Types   []PokemonType `json:"types"`
}

// Similar to location_list.go but retrieves specific area instead of starting from id 0
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		page := Pokemon{}
		err := json.Unmarshal(val, &page)
		if err != nil {
			return Pokemon{}, err
		}

		return page, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil { // if the url failed and couldn't find a location area
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("Error: Pokemon not valid")
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil { // convert the response type to []byte
		return Pokemon{}, err
	}

	pg := Pokemon{}
	err = json.Unmarshal(dat, &pg)
	if err != nil {
		return Pokemon{}, err
	}

	return pg, nil
}
