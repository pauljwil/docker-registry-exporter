package main

import (
	"github.com/pauljwil/docker-registry-exporter/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		logrus.Fatalf("unhandled error: %s", err)
	}
}
