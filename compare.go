// compare 2 string l'une et l'autre
package student

func Compare(a, b string) int {

	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
