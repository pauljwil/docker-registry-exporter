# Docker Registry Exporter 🚀

Export [Docker Registry](https://github.com/distribution/distribution) metrics such as repository and tag counts for [Prometheus](https://prometheus.io/) to scrape. Runs separately from your Docker registry and Prometheus server.

## Build and Run the Application 🛠️


### Docker Build 🐳

Follow these steps to build and run the application using Docker:

1. **Pull the Docker image:**
    ```sh
    docker pull zer0power/docker-registry-exporter:latest
    ```

2. **Run the Docker container:**
    ```sh
    docker run -d -p 9055:9055 -e REGISTRY_ADDRESS=127.0.0.1:5000 zer0power/docker-registry-exporter
    ```
    > **Note:** The `REGISTRY_ADDRESS` environment variable must be set, or the application will crash.

### Manual Build 🪛
Alternatively, you can build and run the Go application in your local environment:

1. **Clone the repository:**
    ```sh
    git clone --depth 1 https://github.com/pauljwil/docker-registry-exporter
    ```

2. **Navigate to the project directory:**
    ```sh
    cd docker-registry-exporter
    ```

3. **Build the application:**
    ```sh
    go build -o docker-registry-exporter
    ```

4. **Run the application:**
    ```sh
    ./docker-registry-exporter
    ```

## Configuration Parameters ⚙️

Configure your Docker Registry Exporter using CLI flags, environment variables, or a configuration file.


| CLI flag | Env var | Config key | Description | Default |
| --- | --- | --- | --- | --- |
| --config | N/A | N/A | Configuration file | docker-registry-exporter.yaml |
| --listen-address | LISTEN_ADDRESS | listen_address | Address to listen on for registry metrics | 127.0.0.1:9055 |
| --metrics-path | METRICS_PATH | metrics_path | Path on which to expose metrics to Prometheus | /metrics |
| --registry-address | REGISTRY_ADDRESS | registry_address | Docker registry address | 127.0.0.1:5000 |

### Example Configuration File 📄
```yaml
listen_address: '127.0.0.1:9055'
metrics_path: '/metrics'
registry_address: '127.0.0.1:5000'
```

## Metrics 📊

The Docker Registry Exporter exposes the following metrics:

| Name | Description | Metric type | Labels |
| --- | --- | --- | --- |
| `repositories` | Number of repositories | Gauge | None |
| `tags` | Number of tags | Gauge | Name |
| `tags_per_repository` | Number of tags per repository | Gauge | `repository` |
| `scrape_latency` | Duration of metrics collection | Gauge | None |
| `scrape_errors` | Number of errors while collecting metrics | Gauge | None |

For more information on gauge values, refer to [Metrics Types](https://prometheus.io/docs/concepts/metric_types/) in the Prometheus documentation.

---

Happy monitoring! 📈✨