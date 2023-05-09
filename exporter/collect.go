package exporter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Repositories struct {
	Repositories []string `json:"repositories"`
}

type Tags struct {
	Tags []string `json:"tags"`
}

func (c *RegistryCollector) Collect(ch chan<- prometheus.Metric) {
	c.countRepositoriesAndTags(ch)
}

func (c *RegistryCollector) countRepositoriesAndTags(ch chan<- prometheus.Metric) {
	totalTags := 0

	repos := c.listRepositories()

	repoCount := len(repos.Repositories)

	ch <- prometheus.MustNewConstMetric(c.metrics.repos, prometheus.GaugeValue, float64(repoCount))

	for _, repo := range repos.Repositories {
		tags := c.listTags(repo)

		tagCount := len(tags.Tags)

		ch <- prometheus.MustNewConstMetric(c.metrics.tagsPerRepo, prometheus.GaugeValue, float64(tagCount), repo)

		totalTags += tagCount
	}

	ch <- prometheus.MustNewConstMetric(c.metrics.tags, prometheus.GaugeValue, float64(totalTags))
}

func (c *RegistryCollector) listRepositories() *Repositories {
	url := "http://" + c.registryAddress + "/v2/_catalog"

	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("failed to get repository list: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("failed to read response body: %s", err)
	}

	repos := &Repositories{}

	json.Unmarshal([]byte(body), repos)

	return repos
}

func (c *RegistryCollector) listTags(repo string) *Tags {
	url := "http://" + c.registryAddress + "/v2/" + repo + "/tags/list"

	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("failed to get tags list: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("failed to read response body: %s", err)
	}

	tags := &Tags{}

	json.Unmarshal([]byte(body), tags)

	return tags
}
