package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type resourceName struct {
	Name string `json:"name"`
}

type page struct {
	Results  []resourceName `json:"results"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
}

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (page, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	// before doing the whole get request process, we see if we already have it
	if val, ok := c.cache.Get(url); ok {
		locationsResp := page{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return page{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil { // if the url failed and couldn't find a location area
		return page{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return page{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil { // convert the response type to []byte
		return page{}, err
	}

	pg := page{}
	err = json.Unmarshal(dat, &pg)
	if err != nil {
		return page{}, err
	}

	c.cache.Add(url, dat) // add it to cache so it can be quickly fetched again
	return pg, nil
}
