package main

import "fmt"

type User struct {
	Name  string
	email string
	age   int
}

func (u *User) PrintUserInfoWithReceiver() {
	fmt.Println("User Name:", u.Name)
	fmt.Println("User Email:", u.email)
	fmt.Println("User Age:", u.age)
}

func (u *User) UpdateUserInfo(name string, email string, age int) {
	u.Name = name
	u.email = email
	u.age = age
	fmt.Print("User Info Updated Successfully!\n")
}

func MethodsAndReceiversExample() {
	user := User{
		Name:  "Alice",
		email: "alice@example.com",
		age:   25,
	}

	user.PrintUserInfoWithReceiver()
	fmt.Print("\n")
	user.UpdateUserInfo("Bob", "bob@example.com", 30)
	fmt.Print("\n")
	user.PrintUserInfoWithReceiver()
}
