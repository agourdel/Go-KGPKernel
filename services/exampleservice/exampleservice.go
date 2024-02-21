package exampleservice

import (
	"agourdel.com/kgpkernel/services/exampleservice/config"
	"agourdel.com/kgpkernel/services/exampleservice/interfaces"
	
)

type ExampleService struct {
	Config config.ExampleConfig
	Provider interfaces.IExampleProvider
}


