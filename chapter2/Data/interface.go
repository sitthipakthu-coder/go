package data

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Person1 struct {
	Name string
}

func (p Person1) Speak() string {
	return "Hello!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Mewe!"
}

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}
