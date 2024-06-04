package main

type Animal interface {
	Says()
	LikesWater() bool
}

type Dog struct{}

func (*Dog) Says() {
	println("Woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

type Cat struct{}

func (*Cat) Says() {
	println("Meow")
}

func (c *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	println(dog.LikesWater())

	cat.Says()
	println(cat.LikesWater())
}

func NewAnimal(factory AnimalFactory) Animal {
	return factory.New()
}
