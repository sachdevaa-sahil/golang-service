package basics

import "fmt"

func CloseConnection() {
	fmt.Println("Connection Closed!")
}

func OpenConnection() {
	fmt.Println("COnnection Opened")
}

func DeferExamples() {
	defer CloseConnection()
	defer MultipleDeferStatementBehaviour()
	defer DeferVariableExample()
	OpenConnection()
}

func MultipleDeferStatementBehaviour() {
	defer fmt.Println("Defer Statement 1")
	defer fmt.Println("Defer Statement 2")
	defer fmt.Println("Defer Statement 3")
	defer fmt.Println("Defer Statement 4")
	defer fmt.Println("Defer Statement 5")
}

func DeferVariableExample() {
	x := 10
	y := 20
	fmt.Println(x + y)
	defer fmt.Println(x + y)

	defer func() {
		fmt.Println("inside the closer func")
		fmt.Println(x + y)
	}()
	x = 40
	y = 60
	fmt.Println(x + y)
}
