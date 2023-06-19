# Docker Registry Exporter

Exports [Docker Registry](https://github.com/distribution/distribution) metrics
such as repository and tag counts for [Prometheus](https://prometheus.io/) to
scrape. Runs separately from your Docker registry and Prometheus server.

## Build and run application

Execute the following commands to build and run the Go application in your
local environment:

```shell script
1. git clone https://github.com/pauljwil/docker-registry-exporter
2. cd docker-registry-exporter
3. go build -o docker-registry-exporter
4. ./docker-registry-exporter
```
## Configuration parameters

Configure your Docker Registry Exporter using either CLI flags, environment
variables, or a configuration file.

| CLI flag | Env var | Config key | Description | Default |
| --- | --- | --- | --- | --- |
| --config | N/A | N/A | Configuration file | docker-registry-exporter.yaml |
| --listen-address | LISTEN_ADDRESS | listen_address | Address to listen on for registry metrics | 127.0.0.1:9055 |
| --metrics-path | METRICS_PATH | metrics_path | Path on which to expose metrics to Prometheus | /metrics |
| --registry-address | REGISTRY_ADDRESS | registry_address | Docker registry address | 127.0.0.1:5000 |

Example configuration file:

```yaml
listen_address: '127.0.0.1:9055'
metrics_path: '/metrics'
registry_address: '127.0.0.1:5000'
```

## Metrics

The Docker Registry Exporter exposes the following metrics:

| Name | Description | Metric type | Labels |
| --- | --- | --- | --- |
| `repositories` | Number of repositories | Gauge | None |
| `tags` | Number of tags | Gauge | Name |
| `tags_per_repository` | Number of tags per repository | Gauge | `repository` |
| `scrape_latency` | Duration of metrics collection | Gauge | None |
| `scrape_errors` | Number of errors while collecting metrics | Gauge | None |

For more information on gauge values, refer to
[Metrics Types](https://prometheus.io/docs/concepts/metric_types/) in the
Prometheus documentation.
