package strategy

import "testing"

func TestStrategy(t *testing.T) {
	var strategy Strategy = ConcreteStrategyA{}
	context := NewContext(strategy)
	context.contextInterface() // StrategyA

	strategy = ConcreteStrategyB{}
	context = NewContext(strategy)
	context.contextInterface() // StrategyB
}
