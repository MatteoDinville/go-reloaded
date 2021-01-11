package utils

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, v := range []rune(s) {
		z01.PrintRune(v)
	}
}
