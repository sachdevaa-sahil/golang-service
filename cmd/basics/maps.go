package main

import "fmt"

func MapsExample() {
	// Create a map to store the ages of people
	users := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	// Print the map
	fmt.Println("Users and their ages:", users)

	for name, age := range users {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Add a new user to the map
	users["David"] = 40
	fmt.Println("\n\n\nAfter adding David:", users)

	// Update an existing user's age
	users["Alice"] = 26
	fmt.Println("After updating Alice's age:", users)

	// Delete a user from the map
	delete(users, "Bob")

	fmt.Printf("After deleting Bob:\n\n\n")

	for name, age := range users {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
