package interfaces

import (
	FooInterfaces "agourdel.com/kgpkernel/domains/foodomain/interfaces"
)

type IRepositoriesManager struct {
	FooRepositories FooInterfaces.IFooRepositories
}