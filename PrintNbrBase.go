/*package main

import (
	"fmt"

	"github.com/01-edu/z01"
	student ".."
)

func main() {
	student.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}*/

package student

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, v := range []rune(s) {
		z01.PrintRune(v)
	}
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		PrintStr("NV")
		return
	}

	if nbr == 0 {
		PrintStr(string(base[0]))
		return
	}

	baseLen := 0
	for range base {
		baseLen++
	}

	if nbr < 0 {
		r := ""
		for nbr != 0 {
			r = string(base[nbr%baseLen*-1]) + r
			nbr /= baseLen
		}
		PrintStr("-" + r)
		return
	}

	r := ""
	for nbr != 0 {
		r = string(base[nbr%baseLen]) + r
		nbr /= baseLen
	}
	PrintStr(r)
}

func isValidBase(s string) bool {
	strLen := 0
	sRune := []rune(s)
	for _, l := range sRune {
		if l == '+' || l == '-' {
			return false
		}
		strLen++
	}
	if strLen < 2 {
		return false
	}
	for i, l := range sRune {
		for j := i + 1; j < strLen; j++ {
			if l == sRune[j] {
				return false
			}
		}
	}
	return true
}
