package basics

import "fmt"

func Operators() {
	// Declare two integer variables
	a := 10
	b := 5

	// Arithmetic Operators
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modulus:", a%b)

	// Relational Operators
	fmt.Println("\nRelational Operators:")
	fmt.Println("Equal to:", a == b)
	fmt.Println("Not equal to:", a != b)
	fmt.Println("Greater than:", a > b)
	fmt.Println("Less than:", a < b)
	fmt.Println("Greater than or equal to:", a >= b)
	fmt.Println("Less than or equal to:", a <= b)

	// Logical Operators
	x := true
	y := false

	fmt.Println("\nLogical Operators:")
	fmt.Println("Logical AND (x && y):", x && y)
	fmt.Println("Logical OR (x || y):", x || y)
	fmt.Println("Logical NOT (!x):", !x)

	// Assignment Operators
	c := 20
	fmt.Println("\nAssignment Operators:")
	fmt.Println("Initial value of c:", c)
	c += 5
	fmt.Println("After c += 5:", c)
	c -= 3
	fmt.Println("After c -= 3:", c)
	c *= 2
	fmt.Println("After c *= 2:", c)
	c /= 4
	fmt.Println("After c /= 4:", c)
	c %= 3
	fmt.Println("After c %= 3:", c)

	// Bitwise Operators
	p := 6 // Binary: 110
	q := 3 // Binary: 011

	fmt.Println("\nBitwise Operators:")
	fmt.Println("Bitwise AND (p & q):", p&q) // Binary: 010 (Decimal: 2)
	fmt.Println("Bitwise OR (p | q):", p|q)  // Binary: 111 (Decimal: 7)
	fmt.Println("Bitwise XOR (p ^ q):", p^q) // Binary: 101 (Decimal: 5)
	fmt.Println("Bitwise AND NOT (p &^ q):", p&^q)

	// increment and decrement operators
	fmt.Println("\nIncrement and Decrement Operators:")
	fmt.Println("Initial value of a:", a)
	a++
	fmt.Println("After incrementing a (a++):", a)

	a--
	fmt.Println("After decrementing a (a--):", a)
}
