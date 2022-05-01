package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (receiver Person) eat() {
	fmt.Println("eat ...", receiver.Name)
}

func main() {
	var p = new(Person)
	p.Id = 1
	p.Name = "chuan"
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(*p)
	fmt.Println(p.Name)

	p.eat()
}
