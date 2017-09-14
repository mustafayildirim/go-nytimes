package nytimes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TopNews(apiKey string) (*Response, error) {
	return getEndpoint("topstories/v2/home.json", apiKey)
}

func getEndpoint(endpoint, apiKey string) (*Response, error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://api.nytimes.com/svc/%s?api-key=%s",
		endpoint, apiKey,
	))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ret := Response{}
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
