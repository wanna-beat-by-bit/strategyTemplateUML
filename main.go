package main

import "fmt"

type Strategy interface {
	Execute()
}

type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Execute() {
	fmt.Println("Executing strategy A")
}

type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Execute() {
	fmt.Println("Executing strategy B")
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	context := Context{}

	strategyA := &ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.ExecuteStrategy()

	strategyB := &ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.ExecuteStrategy()
}
