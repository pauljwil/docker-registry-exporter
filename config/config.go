package config

// Config defines the exporter parameters.
type Config struct {
	// ListenAddress is the HTTP listen endpoint.
	ListenAddress string `mapstructure:"listen_address"`
	// MetricsPath is the HTTP URL path for the registry metrics.
	MetricsPath string `mapstructure:"metrics_path"`

	// RegistryAddress is the host and port for the registry from which metrics
	// will be collected.
	RegistryAddress string `mapstructure:"registry_address"`
}
