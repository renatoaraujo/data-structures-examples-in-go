package linked_lists

import (
	"errors"
	"strconv"
	"strings"
)

// SNode represents a singly linked list node, with only reference to the next node
type SNode struct {
	val  int
	Next *SNode
}

// SinglyLinkedList represents a singly linked list with its head node
type SinglyLinkedList struct {
	head *SNode
}

// NewSinglyLinkedListFromArray creates a singly linked list from an array
func NewSinglyLinkedListFromArray(arr []int) *SinglyLinkedList {
	sll := &SinglyLinkedList{head: nil}
	if len(arr) == 0 {
		return sll
	}

	// get the first element from the array and create the head node
	sll.head = &SNode{val: arr[0]}

	// initialise the previous element to link the other elements, with the first item being the head
	prev := sll.head

	// for each other element it creates a node
	for _, val := range arr[1:] {
		node := &SNode{val: val}
		prev.Next = node
		prev = node
	}

	// returns the singly linked list with the head node
	return sll
}

// DeleteHead deletes the head of the list
func (sll *SinglyLinkedList) DeleteHead() error {
	if sll.head == nil {
		return errors.New("empty list, can't remove first element")
	}

	sll.head = sll.head.Next // Move the head to the next element
	return nil
}

// Prepend add to beginning of the list, setting a new head
func (sll *SinglyLinkedList) Prepend(val int) {
	// creates the new node and set the next to the current head
	n := &SNode{val: val, Next: sll.head}
	// set the head to the new node
	sll.head = n
}

// Append add to the end of the list
func (sll *SinglyLinkedList) Append(val int) {
	n := &SNode{val: val}

	if sll.head == nil {
		sll.head = n
		return
	}

	current := sll.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = n
}

// Delete all matching values from the list
func (sll *SinglyLinkedList) Delete(val int) error {
	if sll.head == nil {
		return errors.New("empty list, can't delete")
	}

	// Remove all occurrences from the head first
	for sll.head != nil && sll.head.val == val {
		sll.head = sll.head.Next
	}

	// If the list became empty after removing the head nodes
	if sll.head == nil {
		return nil
	}

	current := sll.head
	prev := current

	for current != nil {
		if current.val == val {
			// skip the current node to remove it
			prev.Next = current.Next
			// do not update prev if we removed the current
		} else {
			// only update prev if we didn't remove the current
			prev = current
		}
		current = current.Next
	}

	return nil
}

// ToArray converts the singly linked list to array
func (sll *SinglyLinkedList) ToArray() []int {
	if sll.head == nil {
		return []int{}
	}

	var output []int
	current := sll.head
	for current != nil {
		output = append(output, current.val)
		current = current.Next
	}

	return output
}

// Display just returns the string representation of the singly linked list
func (sll *SinglyLinkedList) Display() string {
	if sll.head == nil {
		return "empty list"
	}

	var output []string
	current := sll.head
	for current != nil {
		output = append(output, strconv.Itoa(current.val))
		current = current.Next
	}

	return strings.Join(output, "->")
}
