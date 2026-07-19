package main

import "fmt"

type Persona struct {
	name string
	id   int
	age  int
}

func main() {
	var person1 Persona

	person1.name = "Kaoruko"
	person1.id = 1
	person1.age = 23

	fmt.Println("Name:", person1.name)
	fmt.Println("Id:", person1.id)
	fmt.Println("Age:", person1.age)

}
