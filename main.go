package main

import (
	data "learngo/chapter2/Data"
)

func main() {
	dog := data.Dog{Name: "Buddy"}
	person := data.Person1{Name: "Tang"}
	cat := data.Cat{Name: "Mew"}

	data.MakeSound(dog)
	data.MakeSound(person)
	data.MakeSound(cat)

}
