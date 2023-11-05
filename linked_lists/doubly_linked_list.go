package linked_lists

import (
	"errors"
	"strconv"
	"strings"
)

type DNode struct {
	val  int
	Next *DNode
	Prev *DNode
}

type DoublyLinkedList struct {
	head *DNode
	tail *DNode
}

// NewDoublyLinkedListFromArray creates a singly linked list from an array
func NewDoublyLinkedListFromArray(arr []int) *DoublyLinkedList {
	if len(arr) == 0 {
		return &DoublyLinkedList{head: nil, tail: nil}
	}

	// get the first element from the array and create the head node
	head := &DNode{val: arr[0]}

	// initialise the previous element to link the other elements, with the first item being the head
	last := head

	// for each other element it creates a node
	for _, val := range arr[1:] {
		node := &DNode{val: val, Prev: last}
		last.Next = node
		last = node
	}

	// returns the doubly linked list with the head node
	return &DoublyLinkedList{head: head, tail: last}
}

func (dll *DoublyLinkedList) Head() *DNode {
	return dll.head
}

func (dll *DoublyLinkedList) Tail() *DNode {
	return dll.tail
}

// Prepend add to beginning of the list, setting a new head
func (dll *DoublyLinkedList) Prepend(val int) {
	next := dll.head
	n := &DNode{val: val, Next: next}

	if next == nil {
		dll.head = n
		return
	}
	// creates the new node and set the next to the current head
	// set the prev from the head to the new node
	dll.head.Prev = n
	// set the head to the new node
	dll.head = n
}

// Append adds to the tail of the doubly linked list
func (dll *DoublyLinkedList) Append(val int) {
	n := &DNode{val: val, Prev: dll.tail}

	if dll.tail == nil {
		dll.head = n
	} else {
		dll.tail.Next = n
	}
	dll.tail = n // Update the tail to the new node
}

// DeleteHead deletes the head of the list
func (dll *DoublyLinkedList) DeleteHead() error {
	if dll.head == nil {
		return errors.New("empty list, can't remove head")
	}

	next := dll.head.Next // Store the next element

	if next == nil {
		dll.tail = nil // If there's no next, the list is now empty, and tail should be nil
	} else {
		next.Prev = nil // Otherwise, set the prev of the next element to nil
	}

	dll.head = next // Move the head to the next element

	return nil
}

// Delete all matching values from the list
func (dll *DoublyLinkedList) Delete(val int) error {
	if dll.head == nil {
		return errors.New("empty list, can't delete")
	}

	// Remove all occurrences from the head first
	for dll.head != nil && dll.head.val == val {
		dll.head = dll.head.Next
		if dll.head != nil { // Check if the new head is not nil
			dll.head.Prev = nil // Set the new head's Prev to nil
		}
	}

	// If the list became empty after removing the head nodes
	if dll.head == nil {
		return nil
	}

	current := dll.head

	for current != nil {
		if current.val == val {
			// If it's not the last node, update the next node's Prev
			if current.Next != nil {
				current.Next.Prev = current.Prev
			}
			// If it's not the first node, update the previous node's Next
			if current.Prev != nil {
				current.Prev.Next = current.Next
			}
		}
		current = current.Next
	}

	// After the loop, we need to update the tail if the last node was deleted
	if dll.tail != nil && dll.tail.val == val {
		dll.tail = dll.tail.Prev
		if dll.tail != nil {
			dll.tail.Next = nil
		} else {
			dll.head = nil // The list is now empty
		}
	}

	return nil
}

// DeleteTail deletes the tail of the list
func (dll *DoublyLinkedList) DeleteTail() error {
	if dll.tail == nil {
		return errors.New("empty list, can't remove tail")
	}

	// if there's only one element in the list
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		// set the second-to-last element as the new tail
		dll.tail = dll.tail.Prev
		dll.tail.Next = nil // remove reference to the old tail
	}

	return nil
}

// ToArray converts the doubly linked list to an array
func (dll *DoublyLinkedList) ToArray() []int {
	if dll.head == nil {
		return []int{}
	}

	var output []int
	current := dll.head
	for current != nil {
		output = append(output, current.val)
		current = current.Next
	}

	return output
}

// Display just returns the string representation of the singly linked list
func (dll *DoublyLinkedList) Display() string {
	if dll.head == nil {
		return "empty list"
	}

	var output []string
	current := dll.head
	for current != nil {
		output = append(output, strconv.Itoa(current.val))
		current = current.Next
	}

	return strings.Join(output, "<->")
}
