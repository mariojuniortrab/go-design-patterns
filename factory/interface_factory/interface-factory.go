package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	Name string
	Age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func NewPerson(name string, age int) *person {
	return &person{name, age}
}

func main() {
	p := NewPerson("Jane", 21)
	p.SayHello()
}
