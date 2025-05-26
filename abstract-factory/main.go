package main

import "fmt"

// Animal is a type for our abstract factory example.

type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is a concrete type that implements the Animal interface.

type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("Woof!")
}

func (d *Dog) LikesWater() bool {
	return true
}

// Cat is a concrete type that implements the Animal interface.

type Cat struct{}

func (c *Cat) Says() {
	fmt.Println("Meow!")
}

func (c *Cat) LikesWater() bool {
	return false
}

// AnimalFactory is a abstract factory type for for creating animals.
type AnimalFactory interface {
	New() Animal
}

// DogFactory is a type that implemenets AnimalFactory interface.

type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

// CatFactory is a type that implemenets AnimalFactory interface.

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	dogFactory := &DogFactory{}
	catFactory := &CatFactory{}

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	fmt.Println("Does dog like water?", dog.LikesWater())

	cat.Says()
	fmt.Println("Does cat like water?", cat.LikesWater())
}
