package basics

import "fmt"

func BasicDataTypes() {
	var message string = "Hello, World!"
	var message2 string = "This is a simple Go program!"
	var integer int = 42
	var floater float64 = 3.14
	integer2 := 100        // This is a short variable declaration for an integer
	integer3 := 200        // Another short variable declaration for an integer
	var emptyString string // This is a variable declaration for an empty string
	var emptyInt int       // This is a variable declaration for an empty integer
	var emptyFloat float64 // This is a variable declaration for an empty float
	var emptyBool bool     // This is a variable declaration for an empty boolean
	var a, b, c, d int = 1, 2, 3, 4
	var snake_case_variable string = "This is a snake_case variable"
	var camelCaseVariable string = "This is a camelCase variable"
	var PascalCaseVariable string = "This is a PascalCase variable"
	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(integer)
	fmt.Println(integer2)
	fmt.Println(integer3)
	fmt.Println(floater)
	fmt.Println("The sum of integer2 and integer3 is:", integer2+integer3)

	fmt.Println(emptyBool)
	fmt.Println(emptyString)
	fmt.Println(emptyInt)
	fmt.Println(emptyFloat)
	fmt.Println("The value of a is:", a)
	fmt.Println("The value of b is:", b)
	fmt.Println("The value of c is:", c)
	fmt.Println("The value of d is:", d)
	fmt.Println("Value for PCV CCV SCV are: ", PascalCaseVariable, camelCaseVariable, snake_case_variable)
}
