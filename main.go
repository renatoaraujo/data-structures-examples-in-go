package main

import (
	"fmt"

	"github.com/renatoaraujo/data-structures-examples-in-go/linked_lists"
)

func main() {
	arr := []int{}
	dll := linked_lists.NewDoublyLinkedListFromArray(arr)
	dll.Append(10)
	fmt.Println(dll.Display())
	dll.Append(12)
	fmt.Println(dll.Display())

	err := dll.Delete(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dll.Display())

	dll.Prepend(75)
	fmt.Println(dll.Display())

	err = dll.DeleteHead()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dll.Display())

	dll.Append(12)
	fmt.Println(dll.Display())
	dll.Append(123)
	fmt.Println(dll.Display())

	err = dll.DeleteTail()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dll.Display())

}
