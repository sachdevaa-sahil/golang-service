package basics

import "fmt"

func SingleSwitchCaseExample(day int) {
	fmt.Println("Day of the week for day number", day, ":")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}

func MultipleSwitchCaseExample(day int) {
	fmt.Println("Day of the week for day number month October 2026", day, ":")
	switch day {
	case 1, 8, 15, 22, 29:
		fmt.Println("Thursday")
	case 2, 9, 16, 23, 30:
		fmt.Println("Friday	")
	case 3, 10, 17, 24, 31:
		fmt.Println("Saturday")
	case 4, 11, 18, 25:
		fmt.Println("Sunday")
	case 5, 12, 19, 26:
		fmt.Println("Monday")
	case 6, 13, 20, 27:
		fmt.Println("Tuesday")
	case 7, 14, 21, 28:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}
}
