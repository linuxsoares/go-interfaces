package main

import "fmt"

type Dog interface {
	Bark()
}

type Cat interface {
	Meow()
}

type Shihtzu struct {
	Name string
}

type Chartreux struct {
	Name string
}

func (s Shihtzu) Bark() {
	message := fmt.Sprintf("Type: %T - My name is %s! Bark! Bark! Bark!", s, s.Name)
	fmt.Println(message)
}

func (c Chartreux) Meow() {
	message := fmt.Sprintf("Type: %T - My name is %s! Meow! Meow! Meow! ", c, c.Name)
	fmt.Println(message)
}

func main() {
	var dog Dog
	dog = Shihtzu{Name: "Puppy"}
	dog.Bark()

	var cat Cat
	cat = Chartreux{Name: "Kitty"}
	cat.Meow()
}
