package pointers

import (
	"fmt"
)

func Pointers() {
	var name = "ozgur"

	var calc = func(a int, b int) int {
		if a < 4 {
			panic("a cannot be smaller than 4")
		}
		return a + b
	}(3, 4)

	fmt.Println(&calc)
	fmt.Println(&name)
}
