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

	for i := 0; i < 6; i++ {
		fmt.Println(list.Find(i))
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	list := &LinkedList{}

	for i := 0; i < 6; i++ {
		list.Insert(i, i)
	}
	fmt.Println(list.FindFirst())
	fmt.Println(list.FindLast())
	for i := 0; i < 6; i++ {
		fmt.Println(list.Find(i))
	}

	list.Reverse()
	fmt.Println(list.FindFirst())

	for i := 0; i < 6; i++ {
		fmt.Println(list.Find(i))
	}
}