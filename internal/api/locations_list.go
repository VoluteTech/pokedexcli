package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	locationsResp := RespShallowLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocation{}, err
	}

	return locationsResp, nil
}
