/*package main

import (
	"fmt"
	student ".."
)

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(student.RecursivePower(arg1, arg2))
}*/

package student

func RecursivePower(nb int, power int) int {

	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	}

	if power == 1 {
		return nb
	}
	return nb * RecursivePower(nb, power-1)
}
