package configurable_test

import (
	"fmt"

	"srcd.works/framework/configurable"
)

func ExampleBasicConfiguration() {
	type complexConfiguration struct {
		configurable.BasicConfiguration
		Value string `default:"hola"`
	}

	config := &complexConfiguration{}

	configurable.InitConfig(config)

	fmt.Println(config.Value)
	// Output: hola
}
