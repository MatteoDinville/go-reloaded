/*package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *student.NodeI
	var link2 *student.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(student.SortedListMerge(link2, link))
}*/

package student

func SortedListMerge(n2 *NodeI, n1 *NodeI) *NodeI {
	if n1 == nil && n2 == nil {
		return nil
	} else if n1 == nil && n2 != nil {
		return n2
	} else if n1 != nil && n2 == nil {
		return n1
	}
	l1 := n1
	l2 := n2

	var nouveau *NodeI = nil
	var node *NodeI = nil
	for l1 != nil && l2 != nil {
		if l1.Data > l2.Data {
			if nouveau == nil {
				nouveau = l2
				node = nouveau
			} else {
				nouveau.Next = l2
				nouveau = nouveau.Next
			}
			l2 = l2.Next
		} else {
			if nouveau == nil {
				nouveau = l1
				node = nouveau
			} else {
				nouveau.Next = l1
				nouveau = nouveau.Next
			}
			l1 = l1.Next
		}
	}
	if l1 != nil {
		for l1 != nil {
			nouveau.Next = l1
			nouveau = nouveau.Next
			l1 = l1.Next
		}
	} else if l2 != nil {
		for l2 != nil {
			nouveau.Next = l2
			nouveau = nouveau.Next
			l2 = l2.Next
		}
	}
	return node
}
