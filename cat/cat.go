package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

// print string rune par rune

func printStr(s string) {
	for _, l := range s {
		z01.PrintRune(l)
	}
}

func main() {
	var content []byte
	var tabResult []byte
	var err error

	//si pas d'arguments

	if len(os.Args) == 1 {
		temp := make([]byte, 1)
		var str string
		for {
			os.Stdin.Read(temp)
			str = str + string(temp)
			if temp[0] == '\n' {
				printStr(str)
				str = ""
			}
		}
	}

	// si au moins un argument alors

	//création  d'un tableau avec les caracteres du fichier
	for i := 1; i < len(os.Args); i++ {
		content, err = ioutil.ReadFile(os.Args[i])
		tabResult = append(tabResult, content...)
		if err != nil {
			printStr("cat: ")
			printStr(os.Args[i])
			printStr(": No such file or directory\n")
		}
	}

	//si on a pas d'erreur on print notre tableau, sinon on print l'erreur
	if err == nil {
		printStr(string(tabResult))
	}
}
