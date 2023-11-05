package linked_lists

type SNode struct {
	val  int
	Next *SNode
}

type SinglyLinkedList struct {
	head *SNode
}

func NewSinglyLinkedListFromArray(arr []int) *SinglyLinkedList {
	if len(arr) == 0 {
		return nil
	}

	// get the first element from the array and create the head node
	head := &SNode{val: arr[0]}

	// initialise the previous element to link the other elements, with the first item being the head
	prev := head

	// for each other element it creates a node
	for _, val := range arr[1:] {
		node := &SNode{val: val}
		prev.Next = node
		prev = node
	}

	// returns the singly linked list with the head node
	return &SinglyLinkedList{head: head}
}

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
