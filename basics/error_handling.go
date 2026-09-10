package basics

import (
	"fmt"
)

func Divide(a int, b int) (int, error) {
	if b > 0 {
		return a / b, nil
	} else {
		// return 0, errors.New("Divide by zero")
		return 0, fmt.Errorf("DIVIDE BY ZERO")
	}
}
