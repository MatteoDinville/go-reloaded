/*package main

import student ".."

func main() {
	student.PrintCombN(1)
	student.PrintCombN(3)
	student.PrintCombN(9)
}*/

package student

import "github.com/01-edu/z01"

func ValidComb(nbr int) bool {
	max := 1
	var nbr1 int
	var nbr2 int

	for i := nbr; i > 0; i = (i - i%10) / 10 {
		max *= 10
	}
	max /= 10
	nbr1 = nbr / max
	nbr -= nbr1 * max
	max /= 10
	for i := 1; max > 0; i++ {
		nbr2 = nbr / max
		nbr -= nbr2 * max
		max /= 10
		if nbr1 >= nbr2 {
			return false
		}
		nbr1 = nbr2
	}
	return true
}

func PrintNbr(nbr, n int) {
	count := 0
	max := 1
	var nbrprint int
	z01.PrintRune(',')
	z01.PrintRune(' ')
	for i := nbr; i > 0; i = (i - i%10) / 10 {
		count++
	}
	for i := nbr; i > 0; i = (i - i%10) / 10 {
		max *= 10
	}
	max /= 10
	for i := count; i < n; i++ {
		z01.PrintRune('0')
	}
	for i := 0; i < count; i++ {
		nbrprint = nbr / max
		nbr -= nbrprint * max
		max /= 10
		z01.PrintRune(rune(nbrprint) + '0')
	}
}

func PrintCombN(n int) {
	nbr := 0
	max := 10

	for i := 1; i < n; i++ {
		max *= 10
		max += 10
	}
	for i := 0; i < n; i++ {
		nbr *= 10
		nbr += i
		z01.PrintRune(rune(i) + '0')
	}
	for i := nbr + 1; i < max; i++ {
		if ValidComb(i) == true {
			PrintNbr(i, n)
		}
	}
	z01.PrintRune('\n')
}
