package basics

import "fmt"

func PointersExample() {
	var a int = 10
	var b *int = &a // b is a pointer to a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of b (address of a):", b)
	fmt.Println("Value pointed to by b:", *b)

	// Changing the value of a through the pointer
	*b = 20
	fmt.Println("New value of a after changing through pointer:", a)

	IncrementPointerValue(&a)
	fmt.Println("Value of a after incrementing through pointer:", a)
}

func IncrementPointerValue(p *int) {
	*p++
}

func PrintPersonInfo(person Person) {
	fmt.Println("Person Name:", person.Name)
	fmt.Println("Person Age:", person.Age)
}

func ModifyStructWithoutPointer(person Person) {
	person.Name = "Bob"
	person.Age = 35
	fmt.Println("Modify Struct Without Pointer - Modified Person Name:", person.Name)
	fmt.Println("Modify Struct Without Pointer -Modified Person Age:", person.Age)
}

func ModifyStructWithPointer(person *Person, age int, name string) {
	person.Name = name
	person.Age = age
	fmt.Println("Modify Struct With Pointer - Modified Person Name:", person.Name)
	fmt.Println("Modify Struct With Pointer - Modified Person Age:", person.Age)
}
