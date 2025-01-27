# Embedding config files into go for testability.

This repository shows how to embedd config files into directly in order to:
- run tests to detect config files which are out-of-date,
- print annotated configuration files from application.

It does so by using `go:embed`:

```go
import _ "embed"

//go:embed config.yaml
var exampleConfig string

// ExampleConfig ensures the example config is immutable, since we cannot embed into a const.
func ExampleConfig() string {
    return exampleConfig
}
```

## Testing

The config file is now available in other packages and can be used to test if it matches the current definition:

```go
type appConfig struct {
    Foo string `yaml:"foo"`
}

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
```

[This causing an error is demonstrated in this pull request.](https://github.com/vembacher-unbasical/poc-go-config-files/pull/1)

## Printing Config Files

We now also can easily print the config files, or alternatively write it to a user defined file.

```go
fmt.Printf("%v", examples.ExampleConfig())
```