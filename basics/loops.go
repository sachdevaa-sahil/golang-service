package basics

import "fmt"

func LoopsExample() {
	ForLoopExample()
	WhileLoopExample()
	// InfiniteLoopExample()

}

func ForLoopExample() {
	for i := 0; i <= 10; i++ {
		fmt.Println("The value of i is:", i)
	}
}

func WhileLoopExample() {
	var i int = 0
	for i <= 10 {
		fmt.Println("The value of i is:", i)
		i++
	}
}

func InfiniteLoopExample() {
	for {
		fmt.Println("This is an infinite loop. Press Ctrl+C to stop.")
	}
}
