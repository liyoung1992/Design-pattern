// abstruct_factory project main.go
package main

import (
	"fmt"
)

type IProduct interface {
	show()
}
type Product1 struct {
}
type Product2 struct {
}

func (product *Product1) show() {
	fmt.Println("this is product 1 !!!")
}
func (product *Product2) show() {
	fmt.Println("this is product 2 !!!")
}

type IFactory interface {
	createProduct1()
	createProduct2()
}
type Factory struct {
}

func (factory *Factory) createProduct1() (product1 *Product1) {
	product1 = new(Product1)
	return
}
func (factory *Factory) createProduct2() (product2 *Product2) {
	product2 = new(Product2)
	return
}
func main() {
	factory := new(Factory)
	factory.createProduct1().show()
	factory.createProduct2().show()
	//fmt.Println("Hello World!")
}
