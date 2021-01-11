// package main

// import (
// 	"fmt"
// 	student ".."
// )

// func main() {

// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	student.AdvancedSortWordArr(result, student.Compare)

// 	fmt.Println(result)
// }

package student

func AdvancedSortWordArr(a []string, f func(a, b string) int) {

	//for pour l'index de 0 a l'index max - 1
	for i := 0; i < len(a)-1; i++ {

		// for pour l'index (0 + 1) a l'index max ; sinon on pas de comparaison avec denieres runes
		for j := i + 1; j < len(a); j++ {

			//si notre fonction compare return 1 c'est que a > b
			if f(a[i], a[j]) == 1 {

				//inversion de a et b si a > b
				c := a[i]
				a[i] = a[j]
				a[j] = c
			}
		}
	}
}
