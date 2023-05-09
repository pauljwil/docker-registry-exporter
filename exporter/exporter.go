package exporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RegistryCollector struct {
	listenAddress string
	mux           *http.ServeMux

	registryAddress string

	metrics struct {
		repos       *prometheus.Desc
		tags        *prometheus.Desc
		tagsPerRepo *prometheus.Desc
	}
}

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

func (c *RegistryCollector) ListenAndServe() error {
	serve := http.Server{Addr: c.listenAddress, Handler: c.mux}
	return serve.ListenAndServe()
}
