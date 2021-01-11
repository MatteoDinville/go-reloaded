/*package main

import (
	"fmt"
	student ".."
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Println(student.Split(s, "HA"))
}*/

package student

func Split(s, sep string) []string {

	length := 0
	for i := range s {
		length = i + 1

	}
	len_split := 0
	for i := range sep {
		len_split = i + 1
	}
	for i := 0; i < len_split; i++ {
		s += " "
	}
	prev := false
	len := 0
	for i := 0; i < length; i++ {
		if (s[i:i+len_split] == sep) && !prev {
			prev = true

			len++
		} else {
			prev = false
		}
	}
	len++

	arr := make([]string, len)

	i := 0
	ss := ""
	count := 0
	for i < length {
		if s[i:i+len_split] == sep {
			arr[count] = ss
			count++
			ss = ""
			i += len_split
		}
		ss += string(s[i])
		i += 1
	}

	if ss != "" {
		arr[count] = ss
	}

	return arr

}
