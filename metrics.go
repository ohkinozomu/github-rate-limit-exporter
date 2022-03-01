package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

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
