package themoviedb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	UserAgent   string
	APIEndpoint string
	token       string
	HTTPClient  *http.Client
}

func NewClient() *Client {
	return &Client{
		UserAgent:   getEnv("USER_AGENT", "TV-SERIES-CLI"),
		APIEndpoint: getEnv("API_ENDPOINT", "https://api.themoviedb.org/"),
		token:       getEnv("API_TOKEN_V4", ""),

		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) SarchTV(ctx context.Context, query string, page int, adultContent bool) (*SearchTVResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s3/search/tv?query=%s&page=%d,include_adult=%t", c.APIEndpoint, url.QueryEscape(query), page, adultContent), nil)
	log.Tracef("%s", fmt.Sprintf("%s3/search/tv?query=%s&page=%d,include_adult=%t", c.APIEndpoint, url.QueryEscape(query), page, adultContent))
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	response := SearchTVResponse{}
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetTVSerieSeasons(ctx context.Context, tvID int) (*TVSerieMin, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s3/tv/%d", c.APIEndpoint, tvID), nil)
	log.Tracef("%s", fmt.Sprintf("%s3/tv/%d", c.APIEndpoint, tvID))
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	response := TVSerieMin{}
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetTVSerieEpisodes(ctx context.Context, tvID int, seasonNumber int) (*TVSerieSeasonMin, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s3/tv/%d/season/%d", c.APIEndpoint, tvID, seasonNumber), nil)
	log.Tracef("%s", fmt.Sprintf("%s3/tv/%d/season/%d", c.APIEndpoint, tvID, seasonNumber))
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	response := TVSerieSeasonMin{}
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
