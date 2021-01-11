/*package main

import (
	"fmt"
	student ".."
)

func main() {
	nbits := student.ActiveBits(7)
	fmt.Println(nbits)
}*/

package student

// conversion en binaire et on compte le nombre de "1"

func ActiveBits(n int) uint {
	count := 0
	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return uint(count)
}
