package main

import (
	"fmt"
	"golang-service/users"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()

	fmt.Printf("Main Program: %s \n", id)
	users.GetUsers()
	// basics.DeferExamples()

	// Uncomment an example to run it.
	// basics.BasicDataTypes()
	// basics.ArrayDataTypes()
	// basics.ArrayVsSlice()
	// basics.Operators()
	// basics.SingleSwitchCaseExample(3)
	// basics.MultipleSwitchCaseExample(15)
	// basics.LoopsExample()
	// basics.StructuresExample()
	// basics.MapsExample()
	// basics.PointersExample()
	// var person basics.Person
	// person.Name = "John"
	// person.Age = 30

	// basics.PrintPersonInfo(person)
	// basics.ModifyStructWithoutPointer(person)

	// basics.ModifyStructWithPointer(&person, 40, "Charlie")
	// basics.PrintPersonInfo(person)
	// basics.MethodsAndReceiversExample()
	// basics.InterfacesExample()
	// x, err := basics.Divide(20, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(x)
	// }
}

// var user User
// 	user.Name = "Alice"
// 	// changeUser(&user)

// 	fmt.Println(user.Name) // Output: John
// 	var user2 = &User{Name: "Bob"}
// 	fmt.Println(user2.Name)

// type User struct {
// 	Name string
// }

// func changeUser(u *User) {
// 	u.Name = "John"
// 	u = &User{Name: "Mike"}
// }
