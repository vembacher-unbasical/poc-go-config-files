package main

import (
	"github.com/vembacher-unbasical/poc-go-config-files/examples"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func Test_exampleConfig(t *testing.T) {
	configFile := examples.ExampleConfig()
	var cfg appConfig
	decoder := yaml.NewDecoder(strings.NewReader(configFile))

	// Make sure that unknown fields cause errors.
	decoder.KnownFields(true)
	err := decoder.Decode(&cfg)
	if err != nil {
		t.Fatalf("Error parsing config file: %v", err)
	}
}
