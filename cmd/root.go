package cmd

import (
	"github.com/pauljwil/docker-registry-exporter/config"
	"github.com/pauljwil/docker-registry-exporter/exporter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfg     config.Config
)

var rootCmd = &cobra.Command{
	Use:   "docker-registry-exporter",
	Short: "Prometheus exporter for Docker registry metrics",
	Run: func(cmd *cobra.Command, args []string) {
		collector := exporter.NewRegistryCollector(cfg.ListenAddress, cfg.MetricsPath, cfg.RegistryAddress)

		logrus.Infof("listening on address: %s", cfg.ListenAddress)

		err := collector.ListenAndServe()
		if err != nil {
			logrus.Fatalf("failed to serve http exporter: %s", err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default to docker-registry-exporter.yaml)")
	rootCmd.PersistentFlags().String("listen-address", "127.0.0.1:9055", "Address to listen on for registry metrics")
	rootCmd.PersistentFlags().String("metrics-path", "/metrics", "Path on which to expose metrics to Prometheus")
	rootCmd.PersistentFlags().String("registry-address", "127.0.0.1:5000", "Docker registry address")

	err := viper.BindPFlag("listen_address", rootCmd.PersistentFlags().Lookup("listen-address"))
	if err != nil {
		logrus.Debugf("failed to bind listen address to pflag: %s", err)
	}
	err = viper.BindEnv("listen_address", "LISTEN_ADDRESS")
	if err != nil {
		logrus.Debugf("failed to bind listen address to env variable: %s", err)
	}
	err = viper.BindPFlag("metrics_path", rootCmd.PersistentFlags().Lookup("metrics-path"))
	if err != nil {
		logrus.Debugf("failed to bind metrics path to pflag: %s", err)
	}
	err = viper.BindEnv("metrics_path", "METRICS_PATH")
	if err != nil {
		logrus.Debugf("failed to bind metrics path to env variable: %s", err)
	}
	err = viper.BindPFlag("registry_address", rootCmd.PersistentFlags().Lookup("registry-address"))
	if err != nil {
		logrus.Debugf("failed to bind registry address to pflag: %s", err)
	}
	err = viper.BindEnv("registry_address", "REGISTRY_ADDRESS")
	if err != nil {
		logrus.Debugf("failed to bind registry address to env variable: %s", err)
	}

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("docker-registry-exporter")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.Fatalf("failed to read config file: %s", err)
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		logrus.Fatalf("failed to parse config: %s", err)
	}
}
