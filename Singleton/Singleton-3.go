package main

import (
	"fmt"
	"sync"
)

var m *Manager

var once sync.Once

func GetInstance(name string) *Manager {
	once.Do(func() {
		m = &Manager{name: name}
	})
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