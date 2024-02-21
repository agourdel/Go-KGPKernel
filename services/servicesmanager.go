package services

import(
	KernelInterfaces "agourdel.com/kgpkernel/interfaces"
	ExampleService "agourdel.com/kgpkernel/services/exampleservice"
)

type ServicesManager struct {
	providersManager KernelInterfaces.IProvidersManager
	configsManager KernelInterfaces.IConfigsManager
	ExampleService ExampleService.ExampleService
}

func NewServicesManager(
	configsManager KernelInterfaces.IConfigsManager, 
	providersManager KernelInterfaces.IProvidersManager) *ServicesManager {

		servicesManager := &ServicesManager{
			providersManager: providersManager,
			configsManager: configsManager,
			ExampleService: ExampleService.ExampleService{
				Config:configsManager.ExampleConfig,
				Provider:providersManager.ExampleProvider,
			},
		}

		return servicesManager
}
