package main

import (
	"fmt"

	"github.com/renatoaraujo/data-structures-examples-in-go/linked_lists"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sll := linked_lists.NewSinglyLinkedListFromArray(arr)
	sll.Append(10)

	fmt.Println(sll.Display())
	sll.Append(12)
	fmt.Println(sll.Display())
	sll.Prepend(13)
	fmt.Println(sll.Display())

	err := sll.Delete(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sll.Display())
}
