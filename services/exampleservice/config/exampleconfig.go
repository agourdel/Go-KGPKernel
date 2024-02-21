package config


type ExampleConfig struct {
	FirstValue  int64
	SecondValue string
}

func NewDefautExampleConfig() *ExampleConfig {
	defautConfig := &ExampleConfig{
		FirstValue:  42,
		SecondValue: "Defaut",
	}

	return defautConfig
}

func NewExampleConfig(firstValue int64, secondValue string) *ExampleConfig {
	exampleservice := &ExampleConfig{
		FirstValue: firstValue,
		SecondValue: secondValue,
	}

	return exampleservice
}