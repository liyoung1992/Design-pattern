// Singleton project main.go
package main

import (
	"fmt"
)

var m *Manager

func GetInstance(name string) *Manager {
	if m == nil {
		m = &Manager{name: name}
	}
	return m
}

type Manager struct {
	name string
}

func (this *Manager) say_name() {
	fmt.Println(this.name)
}

func main() {
	man1, man2 := GetInstance("hello"), GetInstance("world")
	man1.say_name()
	man2.say_name()
}
