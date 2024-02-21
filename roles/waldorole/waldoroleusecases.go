package waldoroles

import(
	KernelServices "agourdel.com/kgpkernel/services"
	KernelDomains "agourdel.com/kgpkernel/domains"
	KernelInterfaces "agourdel.com/kgpkernel/interfaces"
)

type WaldoRole struct {
	configManager KernelInterfaces.IConfigsManager
	servicesManager KernelServices.ServicesManager
	domainsManager KernelDomains.DomainsManager
}