package exporter

import "github.com/prometheus/client_golang/prometheus"

func (c *RegistryCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *RegistryCollector) initMetrics() {
	// Registry metrics
	c.metrics.repos = prometheus.NewDesc(
		"repositories",
		"Number of repositories",
		nil, nil,
	)
	c.metrics.tags = prometheus.NewDesc(
		"tags",
		"Number of tags",
		nil, nil,
	)
	c.metrics.tagsPerRepo = prometheus.NewDesc(
		"tags_per_repository",
		"Number of tags per repository",
		[]string{"repository"}, nil,
	)

	// Prometheus scrape metrics
	c.metrics.scrapeLatency = prometheus.NewDesc(
		"scrape_latency",
		"Duration of metrics collection",
		nil, nil,
	)
	c.metrics.scrapeErrors = prometheus.NewDesc(
		"scrape_errors",
		"Number of errors while collecting metrics",
		nil, nil,
	)
}
