package main

import (
	"fmt"
	"github.com/vembacher-unbasical/poc-go-config-files/examples"
)

type appConfig struct {
	Foo string `yaml:"foo"`
}

func main() {
	fmt.Printf("%v", examples.ExampleConfig())
}
