package interfaces 

import(
	ExampleInterfaces "agourdel.com/kgpkernel/services/exampleservice/interfaces"
)

type IProvidersManager struct {
	ExampleProvider ExampleInterfaces.IExampleProvider
}