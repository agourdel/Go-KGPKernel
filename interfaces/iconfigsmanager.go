package interfaces 

import(
	FooConfig "agourdel.com/kgpkernel/domains/foodomain/config"
	ExampleConfig "agourdel.com/kgpkernel/services/exampleservice/config"
)

type IConfigsManager struct {
	FooConfig FooConfig.FooConfig
	ExampleConfig ExampleConfig.ExampleConfig
}