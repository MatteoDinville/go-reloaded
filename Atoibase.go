/*package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("1111101", "01"))
	fmt.Println(student.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
}*/

package student

// atoi pour passer en int et ensuite conversion de base
// return base converti a la base une string

func AtoiBase(s string, base string) int {
	finalint := 0

	lenbase := 0
	for i := range base {
		lenbase = i + 1
	}

	if !isbasevalid(base, lenbase) {
		return finalint
	}

	lens := 0
	for i := range s {
		lens = i + 1
	}

	for i := range s {
		for j := range base {
			if s[i] == base[j] {
				finalint = finalint + j
				if i != lens-1 {
					finalint = finalint * lenbase
				}
			}
		}
	}

	return finalint
}

func isbasevalid(base string, lenbase int) bool {
	if lenbase < 2 {
		return false
	}

	for i := range base {
		for j := range base {
			if base[i] == base[j] && i != j {
				return false
			}
		}
	}

	for i := range base {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
	}

	return true
}
