package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Repositories struct {
	Repositories []string `json:"repositories"`
}

type Tags struct {
	Tags []string `json:"tags"`
}

// Collect sends collected metrics to the Prometheus channel.
func (c *RegistryCollector) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()

	errcount := c.countRepositoriesAndTags(ch)

	elapsed := time.Since(start)

	ch <- prometheus.MustNewConstMetric(c.metrics.scrapeErrors, prometheus.GaugeValue, float64(errcount))
	ch <- prometheus.MustNewConstMetric(c.metrics.scrapeLatency, prometheus.GaugeValue, elapsed.Seconds())
}

// countRepositoriesAndTags counts the number of repositories, the number of
// tags per repository, and the total number of tags and sends the counts as
// metrics to the Prometheus channel.
func (c *RegistryCollector) countRepositoriesAndTags(ch chan<- prometheus.Metric) (errcount int) {
	errcount = 0
	totalTags := 0

	repos, err := c.listRepositories()
	if err != nil {
		logrus.Errorf("failed to compile repository list: %s", err)
		errcount++
	}

	repoCount := len(repos.Repositories)

	ch <- prometheus.MustNewConstMetric(c.metrics.repos, prometheus.GaugeValue, float64(repoCount))

	for _, repo := range repos.Repositories {
		tags, err := c.listTags(repo)
		if err != nil {
			logrus.Errorf("failed to compile tags list: %s", err)
			errcount++
		}

		tagCount := len(tags.Tags)

		ch <- prometheus.MustNewConstMetric(c.metrics.tagsPerRepo, prometheus.GaugeValue, float64(tagCount), repo)

		totalTags += tagCount
	}

	ch <- prometheus.MustNewConstMetric(c.metrics.tags, prometheus.GaugeValue, float64(totalTags))

	return errcount
}

// listRepositories returns a list of the repositories present in the target
// registry.
func (c *RegistryCollector) listRepositories() (*Repositories, error) {
	url := ""

	if strings.HasPrefix(c.registryAddress, "http://") {
		url = c.registryAddress + "/v2/_catalog"
	} else {
		url = "http://" + c.registryAddress + "/v2/_catalog"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query repository list: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	repos := &Repositories{}

	json.Unmarshal([]byte(body), repos)

	return repos, nil
}

// listTags returns a list of the tags present in a given repository.
func (c *RegistryCollector) listTags(repo string) (*Tags, error) {
	url := ""

	if strings.HasPrefix(c.registryAddress, "http://") {
		url = c.registryAddress + "/v2/" + repo + "/tags/list"
	} else {
		url = "http://" + c.registryAddress + "/v2/" + repo + "/tags/list"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query tags list: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}

	tags := &Tags{}

	json.Unmarshal([]byte(body), tags)

	return tags, nil
}
