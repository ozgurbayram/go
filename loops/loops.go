package loops

import (
	"fmt"
	"mygogo/functions"
)

func Loop() {
	for i := 0; i < 30; i++ {
		total, diff := functions.Fn()
		fmt.Println("total", total, "diff", diff)
	}
}
