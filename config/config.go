package config

type Config struct {
	ListenAddress string `mapstructure:"listen_address"`
	MetricsPath   string `mapstructure:"metrics_path"`

	RegistryAddress string `mapstructure:"registry_address"`
}
