package list

import (
	"testing"
	"fmt"
)


func TestLinkedList_Insert(t *testing.T) {
	list := &LinkedList{}

	for i := 0; i < 6; i++ {
		list.Insert(i, i)
	}
	fmt.Println(list.FindFirst())
	fmt.Println(list.FindLast())
	list.PrintList()
}

func TestLinkedList_Reverse(t *testing.T) {
	list := &LinkedList{}

	for i := 0; i < 6; i++ {
		list.Insert(i, i)
	}
	fmt.Println(list.FindFirst())
	fmt.Println(list.FindLast())
	list.PrintList()

	list.Reverse()
	fmt.Println(list.FindFirst())

	//list.PrintList()
}