package main

import "fmt"

func ArrayDataTypes() {
	// Declare an array of integers with a fixed size of 5
	var intArray [5]int = [5]int{1, 2, 3, 4, 5}

	// Declare an array of strings with a fixed size of 3
	var stringArray [3]string = [3]string{"Go", "is", "awesome"}

	// Accessing elements in the integer array
	fmt.Println("Integer Array Elements:")
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("Element at index %d: %d\n", i, intArray[i])
	}

	// Accessing elements in the string array
	fmt.Println("\nString Array Elements:")
	for i := 0; i < len(stringArray); i++ {
		fmt.Printf("Element at index %d: %s\n", i, stringArray[i])
	}

	// Modifying elements in the integer array
	intArray[0] = 10
	intArray[1] = 20

	fmt.Println("\nModified Integer Array Elements:")
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("Element at index %d: %d\n", i, intArray[i])
	}

	// Using the built-in copy function to copy elements from one array to another
	var newIntArray [5]int
	copy(newIntArray[:], intArray[:])

	fmt.Println("\nCopied Integer Array Elements:")
	for i := 0; i < len(newIntArray); i++ {
		fmt.Printf("Element at index %d: %d\n", i, newIntArray[i])
	}
}

func ArrayVsSlice() {
	// Declare an array of integers with a fixed size of 5
	var intArray [5]int = [5]int{1, 2, 3, 4, 5}

	// Declare a slice of integers
	intSlice := []int{10, 20, 30, 40, 50}

	// Accessing elements in the integer array
	fmt.Println("Integer Array Elements:")
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("Element at index %d: %d\n", i, intArray[i])
	}

	// Accessing elements in the integer slice
	fmt.Println("\nInteger Slice Elements:")
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("Element at index %d: %d\n", i, intSlice[i])
	}

	// Modifying elements in the integer slice
	intSlice[0] = 100
	intSlice[1] = 200

	fmt.Println("\nModified Integer Slice Elements:")
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("Element at index %d: %d\n", i, intSlice[i])
	}

	// Using the built-in append function to add elements to the slice
	intSlice = append(intSlice, 60, 70)

	fmt.Println("\n\nlength of intSlice:", len(intSlice))
	fmt.Println("\nAppended Integer Slice Elements:")
	for i := 0; i < len(intSlice); i++ {
		fmt.Printf("Element at index %d: %d\n", i, intSlice[i])
	}
}
