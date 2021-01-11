/*package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.Atoi("12345"))
	fmt.Println(student.Atoi("0000000012345"))
	fmt.Println(student.Atoi("012 345"))
	fmt.Println(student.Atoi("Hello World!"))
	fmt.Println(student.Atoi("+1234"))
	fmt.Println(student.Atoi("-1234"))
	fmt.Println(student.Atoi("++1234"))
	fmt.Println(student.Atoi("--1234"))
}*/

package student

//converti les string en entier (int) "string en int"

func Atoi(s string) int {
	result := 0
	carac := s
	signe := 1

	if s == "" {
		return 0
	}
	if s[0] == '-' {

		carac = s[1:]
		signe = -1

	} else if s[0] == '+' {

		carac = s[1:]

	}
	for _, cha := range carac {

		if cha >= '0' && cha <= '9' {

			chiffre := int(cha) - 48
			result = result*10 + chiffre

		} else {

			return 0
		}
	}

	return signe * result
}
