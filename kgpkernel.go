package kgpkernel

import(
	WaldoRole "agourdel.com/kgpkernel/roles/waldorole"
	KernelInterfaces "agourdel.com/kgpkernel/interfaces"
	KernelServices "agourdel.com/kgpkernel/services"
	KernelDomains "agourdel.com/kgpkernel/domains"
)

type Kernel struct {
	RepositoriesManager KernelInterfaces.IRepositoriesManager
	ProvidersManager KernelInterfaces.IProvidersManager
	ConfigManager KernelInterfaces.IConfigsManager
	ServicesManager KernelServices.ServicesManager
	DomainesManager KernelDomains.DomainsManager
	Roles RolesStruct
}


type RolesStruct struct {
	WaldoRole WaldoRole.WaldoRole 
}


func newKernel(
	config KernelInterfaces.IConfigsManager, 
	repositoriesManager KernelInterfaces.IRepositoriesManager,
	providersManager KernelInterfaces.IProvidersManager) *Kernel {

		kernel := &Kernel{
			RepositoriesManager : repositoriesManager,
			ProvidersManager: providersManager,
			ConfigManager: config,
			ServicesManager: *KernelServices.NewServicesManager(config,providersManager),
		}

		return kernel
	
	}