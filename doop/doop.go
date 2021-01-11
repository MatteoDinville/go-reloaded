package main

import (
	"os"

	student ".."
	"github.com/01-edu/z01"
)

func PrintNbr(nbr int64) {
	var count int64
	var nbrprint int64
	count = 1

	if nbr < 0 {
		nbr *= -1
		z01.PrintRune('-')
	}
	for i := nbr; i > 9; i /= 10 {
		count *= 10
	}
	for count >= 1 {
		nbrprint = nbr / count
		nbr = nbr - nbrprint*count
		count /= 10
		z01.PrintRune(rune(nbrprint) + '0')
	}
	z01.PrintRune('\n')
}

func PrintStr(s string) {
	for _, l := range s {
		z01.PrintRune(l)
	}
}

// si il y a que des chiffres dans tt les args
func DoopError(s string) bool {
	if len(s) > 1 {
		return false
	}
	if s[0] != '+' && s[0] != '-' && s[0] != '*' && s[0] != '/' && s[0] != '%' {
		return false
	}
	for _, l := range os.Args[1] {
		if l < '0' || l > '9' {
			return false
		}
	}
	for _, l := range os.Args[3] {
		if l < '0' || l > '9' {
			return false
		}
	}
	return true
}

func main() {
	if DoopError(os.Args[2]) == false {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	op1 := student.Atoi(os.Args[1])
	op2 := student.Atoi(os.Args[3])
	if op1 >= 9223372036854775807 || op1 <= -9223372036854775807 || op2 >= 9223372036854775807 || op2 <= -9223372036854775807 {
		PrintStr("0\n")
		return
	}
	if os.Args[2][0] == '/' && op2 == 0 {
		PrintStr("No division by 0\n")
		return
	}
	if os.Args[2][0] == '%' && op2 == 0 {
		PrintStr("No modulo by 0\n")
		return
	}
	if os.Args[2][0] == '+' {
		op1 += op2
	}
	if os.Args[2][0] == '-' {
		op1 -= op2
	}
	if os.Args[2][0] == '*' {
		op1 *= op2
	}
	if os.Args[2][0] == '/' {
		op1 /= op2
	}
	if os.Args[2][0] == '%' {
		op1 %= op2
	}
	PrintNbr(int64(op1))
}
