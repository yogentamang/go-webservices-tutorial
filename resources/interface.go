package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
