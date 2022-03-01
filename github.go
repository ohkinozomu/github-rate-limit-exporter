package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
)

type Response struct {
	Resources struct {
		Core struct {
			Limit     int `json:"limit"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
			Used      int `json:"used"`
		} `json:"core"`

		GraphQL struct {
			Limit     int `json:"limit"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
			Used      int `json:"used"`
		} `json:"graphql"`

		IntegrationManifest struct {
			Limit     int `json:"limit"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
			Used      int `json:"used"`
		} `json:"integration_manifest"`

		Search struct {
			Limit     int `json:"limit"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
			Used      int `json:"used"`
		} `json:"search"`
	} `json:"resources"`

	Rate struct {
		Limit     int `json:"limit"`
		Remaining int `json:"remaining"`
		Reset     int `json:"reset"`
		Used      int `json:"used"`
	} `json:"rate"`
}

func getRateLimit() (Response, error) {
	var r Response
	ctx := context.Background()
	sts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	client := oauth2.NewClient(ctx, sts)
	url := "https://api.github.com/rate_limit"

	resp, err := client.Get(url)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	}
	return r, nil
}
