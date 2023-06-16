package exporter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/require"
)

var (
	listenAddress string = "127.0.0.1:9055"
	metricsPath   string = "/metrics"
)

func TestNewRegistryCollector(t *testing.T) {
	reposResponseBody := `{"repositories":["my-repo"]}`
	tagResponseBody := `{"name":"my-repo","tags":["latest"]}`

	// mockServer mocks HTTP responses for each of the API calls made when
	// collecting metrics.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v2/_catalog":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			responseBody := []byte(reposResponseBody)
			w.Write(responseBody)
		case "/v2/my-repo/tags/list":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			responseBody := []byte(tagResponseBody)
			w.Write(responseBody)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	registry := prometheus.NewRegistry()

	collector := NewRegistryCollector(listenAddress, metricsPath, mockServer.URL)

	registry.MustRegister(collector)

	metricsHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	metricsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricsHandler.ServeHTTP(w, r)
	}))
	defer metricsServer.Close()

	resp, err := http.Get(metricsServer.URL)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	expectedMetrics := []string{
		"repos",
		"tags",
		"tags_per_repository",
		"scrape_latency",
		"scrape_errors",
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	bodyString := string(bodyBytes)

	for _, metricName := range expectedMetrics {
		require.Contains(t, bodyString, metricName)
	}
}
