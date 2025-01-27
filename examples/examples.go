package examples

import _ "embed"

//go:embed config.yaml
var exampleConfig string

// ExampleConfig ensures the example config is immutable, since we cannot embed into a const.
func ExampleConfig() string {
	return exampleConfig
}
