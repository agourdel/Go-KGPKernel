package foodomain

import (
	"agourdel.com/kgpkernel/domains/foodomain/config"
	"agourdel.com/kgpkernel/domains/foodomain/interfaces"


)

type FooActions struct {
	fooConfig       config.FooConfig
	fooRepositories interfaces.IFooRepositories
	fooService1     interfaces.IFooService1
}

func newFooActions(fooConfig config.FooConfig, fooRepositories interfaces.IFooRepositories, fooService1 interfaces.IFooService1) *FooActions {
	mFooActions := &FooActions{
		fooConfig:       fooConfig,
		fooRepositories: fooRepositories,
		fooService1:     fooService1,
	}

	return mFooActions
}
