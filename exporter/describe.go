package exporter

import "github.com/prometheus/client_golang/prometheus"

func (c *RegistryCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *RegistryCollector) initMetrics() {
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
}
