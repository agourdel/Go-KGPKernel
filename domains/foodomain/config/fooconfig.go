package config

type FooConfig struct {
	FirstValue  int64
	SecondValue string
}

func NewDefautFooConfig() *FooConfig {
	defautConfig := &FooConfig{
		FirstValue:  42,
		SecondValue: "Defaut",
	}

	return defautConfig
}

func NewFooConfig(firstValue int64, secondValue string) *FooConfig {
	fooConfig := &FooConfig{
		FirstValue:  firstValue,
		SecondValue: secondValue,
	}

	return fooConfig
}
