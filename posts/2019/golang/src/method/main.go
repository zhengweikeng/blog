package main

import "fmt"

type Animal interface {
	Bark()
}

type Dog struct {
	Name string
}

func (dog Dog) SetName(name string) {
	dog.Name = name
}

func (dog Dog) Bark() {
	fmt.Println("Wang")
}

type Cat struct {
	Name string
}

func (cat *Cat) SetName(name string) {
	cat.Name = name
}

func (cat *Cat) Bark() {
	fmt.Println("Miao")
}

func GetCat(name string) Cat {
	return Cat{name}
}

func Bark(animal Animal) {
	animal.Bark()
}

func main() {
	dog := Dog{"little dog"}
	dog.SetName("wangcai")
	fmt.Println(dog.Name) // little dog

	cat := &Cat{"little cat"}
	cat.SetName("Wangcai")
	fmt.Println(cat.Name) // Wangcai

	dog2 := &Dog{"little dog"}
	dog2.SetName("wangcai")
	// (*dog2).SetName("wangcai")
	fmt.Println(dog2.Name) // little dog

	cat2 := Cat{"little cat"}
	cat2.SetName("Wangcai")
	// (&cat2).SetName("Wangcai")
	fmt.Println(cat2.Name) // Wangcai

	// GetCat("little Cat").SetName("Wangcai")

	var animal Animal = dog
	animal.Bark()

	var animal2 Animal = &Cat{}
	animal2.Bark()

	// var animal3 Animal = Cat{"little cat"}
	// animal.Bark()
}
