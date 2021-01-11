package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	} else {
		var strArg string
		var tabResult []rune
		var tabIndex []int
		var tabVowels []rune
		var tabVowelsReverse []rune

		//creer une string avec les arguments
		for indexStr, str := range os.Args[1:] {
			if indexStr != 0 {
				strArg = strArg + " " + str
			} else {
				strArg = strArg + str
			}

		}

		//creer un tableau d'index de voyelles et un tableau de voyelles
		for index, letter := range strArg {
			if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' {
				tabIndex = append(tabIndex, index)
				tabVowels = append(tabVowels, letter)
			}
		}

		//créer un tableau avec les voyelles inversées
		for i := len(tabVowels) - 1; i >= 0; i-- {
			tabVowelsReverse = append(tabVowelsReverse, tabVowels[i])
		}

		//remplacer les voyelles par leur oppose present dans le tableau
		tabResult = []rune(strArg)

		for i := 0; i < len(tabVowelsReverse); i++ {
			tabResult[tabIndex[i]] = tabVowelsReverse[i]
		}
		fmt.Println(string(tabResult))
	}
}
