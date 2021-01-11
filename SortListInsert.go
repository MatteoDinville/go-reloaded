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

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = student.SortListInsert(link, -2)
	link = student.SortListInsert(link, 2)
	PrintList(link)
}*/

package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return nil
	}
	node := l
	n := &NodeI{Data: data_ref}
	var previous *NodeI = nil
	for l.Next != nil {
		if l.Data > data_ref {
			if previous == nil {
				node = n
				n.Next = l
			} else {
				previous.Next = n
				n.Next = l
			}
			return node
		}
		previous = l
		l = l.Next
	}
	if l.Data > data_ref {
		previous.Next = n
		n.Next = l
	} else {
		l.Next = n
	}
	return node
}
