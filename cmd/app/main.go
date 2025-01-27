package main

import (
	"fmt"
	"github.com/vembacher-unbasical/poc-go-config-files/examples"
)

type appConfig struct {
	Bar string `yaml:"bar"`
}

func main() {
	fmt.Printf("%v", examples.ExampleConfig())
}
