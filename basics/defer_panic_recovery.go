package basics

import "fmt"

func PanicWithDeferExample() {
	fmt.Println("program Start!")

	defer fmt.Println("Program ShutDown Properly")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered!")
		}
	}()
	fmt.Println("Program in between execution")

	if true {
		panic("SOMETHING HAS WENT WRONG")
	}

	fmt.Println("PROGRAM EXECUTION COMPLETE")
}
