// factory project main.go
package main

import (
	"fmt"
)

type BaseOperation struct {
	num1 float64
	num2 float64
}
type OperationAdd struct {
	BaseOperation
}
type OperationSub struct {
	BaseOperation
}
type Operation interface {
	getResult() float64
	setNumA(float64)
	setNumB(float64)
}

//-----------------------
type Factory interface {
	createOperation() Operation
}

type AddFactory struct {
}

func (a *AddFactory) createOperation() (operation Operation) {
	operation = new(OperationAdd)
	return
}

//-----------------------
type SubFactory struct {
}

func (s *SubFactory) createOperation() (operation Operation) {
	operation = new(OperationSub)
	return
}

//-----------------------
func (operation *BaseOperation) setNumA(numA float64) {
	operation.num1 = numA
}
func (operation *BaseOperation) setNumB(numB float64) {
	operation.num2 = numB
}
func (operation *OperationAdd) getResult() float64 {
	return operation.num1 + operation.num2
}
func (operation *OperationSub) getResult() float64 {
	return operation.num1 - operation.num2
}

//-----------------------
func main() {
	add_factory := new(AddFactory)
	add_oper := add_factory.createOperation()
	add_oper.setNumA(10)
	add_oper.setNumB(20)
	fmt.Println(add_oper.getResult())
	//	fmt.Println("Hello World!")
}
