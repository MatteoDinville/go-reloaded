package student

import (
	"github.com/01-edu/z01"
	utils."./utils"
)

func PrintNbr(n int) {
	if n == 0 {

		utils.PrintStr("0/n")
		return
	} else if n == -9223372036854775808 {
		utils.PrintStr("-9223372036854775808\n")
		return

	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	nbDec := 0
	for k := n; k > 0; k /= 10 {
		nbDec++
	}
	s := make([]rune, nbDec)
	for l := n; l > 0; l /= 10 {
		s[nbDec-1] = rune('0' + l%10)
	}
	utils.PrintStr(string(s))
	z01.PrintRune('\n')
}
