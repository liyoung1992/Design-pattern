// sample_factory project main.go
package main

import (
	"fmt"
)

type Operation interface {
	getResult() float64
	setNumA(float64)
	setNumB(float64)
}

type BaseOperation struct {
	numberA float64
	numberB float64
}

func (Operation *BaseOperation) setNumA(numA float64) {
	Operation.numberA = numA
}
func (Operation *BaseOperation) setNumB(numB float64) {
	Operation.numberB = numB
}

type OperationAdd struct {
	BaseOperation
}

func (this *OperationAdd) getResult() float64 {
	return this.numberA + this.numberB
}

type OperationSub struct {
	BaseOperation
}

func (this *OperationSub) getResult() float64 {
	return this.numberA - this.numberB
}

type OperationFactory struct {
}

func (this OperationFactory) createOperation(operator string) (operation Operation) {
	switch operator {
	case "+":
		operation = new(OperationAdd)
	case "-":
		operation = new(OperationSub)
	default:
		panic("error")
	}
	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var fac OperationFactory
	oper := fac.createOperation("+")
	oper.setNumA(4.0)
	oper.setNumB(5.0)
	fmt.Println(oper.getResult())
}
