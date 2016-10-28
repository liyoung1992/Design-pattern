// observer project main.go
package main

import (
	"container/list"
	"fmt"
)

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}
type Observer interface {
	update(Subject)
}

type ConcreteSubject struct {
	observers *list.List
	value     int
}

func NewConcreteSubject() *ConcreteSubject {
	s := new(ConcreteSubject)
	s.observers = list.New()
	return s
}
func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers.PushBack(observer)
}
func (s *ConcreteSubject) Detach(observer Observer) {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(*Observer) == &observer {
			s.observers.Remove(ob)
			break
		}
	}
}
func (s *ConcreteSubject) Notify() {
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(Observer).update(s)
	}
}
func (s *ConcreteSubject) setValue(v int) {
	s.value = v
	s.Notify()
}
func (s *ConcreteSubject) getValue() (value int) {
	return s.value
}

type ConcreteObserver1 struct {
}

func (o *ConcreteObserver1) update(subject Subject) {
	fmt.Println("observser1 value is ", subject.(*ConcreteSubject).getValue())
}

type ConcreteObserver2 struct {
}

func (o *ConcreteObserver2) update(subject Subject) {
	fmt.Println("observser2 value is ", subject.(*ConcreteSubject).getValue())
}

func main() {
	subject := NewConcreteSubject()
	observer1 := new(ConcreteObserver1)
	observer2 := new(ConcreteObserver2)
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.setValue(55)
	//	fmt.Println("Hello World!")
}
