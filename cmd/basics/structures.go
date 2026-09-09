package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func StructuresExample() {
	person := Person{
		Name: "Alice",
		Age:  25,
	}

	fmt.Println("Person Name:", person.Name)
	fmt.Println("Person Age:", person.Age)
}
