package exporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegistryCollector is a Prometheus collector for the Docker registry metrics.
type RegistryCollector struct {
	listenAddress string
	mux           *http.ServeMux

	registryAddress string

	metrics struct {
		repos       *prometheus.Desc
		tags        *prometheus.Desc
		tagsPerRepo *prometheus.Desc

		scrapeLatency *prometheus.Desc
		scrapeErrors  *prometheus.Desc
	}
}

// NewRegistryCollector creates a new RegistryCollector, registers its metrics
// with the default Prometheus Registerer, and configures the handler for the
// metrics endpoint.
func NewRegistryCollector(listenAddress, metricsPath, registryAddress string) *RegistryCollector {
	collector := &RegistryCollector{
		listenAddress:   listenAddress,
		registryAddress: registryAddress,
	}

	collector.initMetrics()

	prometheus.MustRegister(collector)

	collector.mux = http.NewServeMux()
	collector.mux.Handle(
		metricsPath,
		promhttp.InstrumentMetricHandler(
			prometheus.DefaultRegisterer,
			promhttp.HandlerFor(prometheus.DefaultGatherer,
				promhttp.HandlerOpts{},
			),
		),
	)

	return collector
}

// ListenAndServe creates the Prometheus HTTP server that exports the Docker
// registry metrics for Prometheus to scrape.
func (c *RegistryCollector) ListenAndServe() error {
	serve := http.Server{Addr: c.listenAddress, Handler: c.mux}
	return serve.ListenAndServe()
}
