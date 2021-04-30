package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

type Metrics struct {
	ResourcesCoreRemaining                prometheus.Gauge
	ResourcesGraphQLRemaining             prometheus.Gauge
	ResourcesIntegrationManifestRemaining prometheus.Gauge
	ResourcesSearchRemaining              prometheus.Gauge
}

func NewMetrics() *Metrics {
	rcr := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_rate_limit_exporter_resources_core_remaining",
		Help: "Core Remaining",
	})
	rgr := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_rate_limit_exporter_resources_graphql_remaining",
		Help: "GraphQL Remaining",
	})
	rir := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_rate_limit_exporter_resources_integrationmanifest_remaining",
		Help: "IntegrationManifest Remaining",
	})
	rsr := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_rate_limit_exporter_resources_search_remaining",
		Help: "Search Remaining",
	})
	return &Metrics{
		ResourcesCoreRemaining:                rcr,
		ResourcesGraphQLRemaining:             rgr,
		ResourcesIntegrationManifestRemaining: rir,
		ResourcesSearchRemaining:              rsr,
	}
}

func (m *Metrics) record() {
	for {
		r, err := getRateLimit()
		if err != nil {
			log.Println(err)
		}
		m.ResourcesCoreRemaining.Set(float64(r.Resources.Core.Remaining))
		m.ResourcesGraphQLRemaining.Set(float64(r.Resources.GraphQL.Remaining))
		m.ResourcesIntegrationManifestRemaining.Set(float64(r.Resources.IntegrationManifest.Remaining))
		m.ResourcesSearchRemaining.Set(float64(r.Resources.Search.Remaining))
		time.Sleep(10 * time.Second)
	}
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

func main() {
	m := NewMetrics()
	go m.record()

	go runFeatures()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
