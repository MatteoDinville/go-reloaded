/*package main

import (
	"fmt"
	student ".."
)

func main() {
	fmt.Println(student.SplitWhiteSpaces("Hello how are you?"))
}*/

package student

// lorsqu'il y un espace ca ajoute dans une case d'un tableau

func SplitWhiteSpaces(s string) []string {
	count := 0
	size := len(s) - 1

	for i, _ := range s {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if i < size && s[i+1] != ' ' && s[i+1] != '\t' && s[i+1] != '\n' {
				count++
			}
		}
	}
	tab := make([]string, count+1)
	count = 0

	for i, _ := range s {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if i < size && s[i+1] != ' ' && s[i+1] != '\t' && s[i+1] != '\n' {
				count++
			}
		} else {
			tab[count] += string(s[i])
		}
	}
	return tab
}
