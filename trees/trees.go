package trees

import (
	"fmt"
	"sort"
)

// Node is the representation of a node, containing its data (in this case an integer) and the left and right nodes references
type Node struct {
	Data  int
	left  *Node
	right *Node
}

// NewBalancedBSTFromSlice will create a tree from an array of integer, by splitting it in 2 and inserting the nodes in their
// respective positions using the recursive Insert method
func NewBalancedBSTFromSlice(arr []int) *Node {
	l := len(arr)
	if l == 0 {
		return nil
	}

	// to an effective balanced BST tree the array should be sorted
	sort.Ints(arr)

	// get the middle of the array to split in 2 pieces
	mid := l / 2
	root := &Node{Data: arr[mid]}

	// iterate over the first half of the array and insert the nodes
	for _, val := range arr[:mid] {
		Insert(root, val)
	}

	// iterate over the second half of the array and insert the nodes
	for _, val := range arr[mid+1:] {
		Insert(root, val)
	}

	return root
}

// Insert will add the val to the tree, by recursively looking through the values in the left and the right side
func Insert(root *Node, val int) *Node {
	// root node is not set, so we create a new one here, by giving the current val
	if root == nil {
		return &Node{Data: val}
	} else {
		// check if the val is greater than the current val, in which will recursively insert into the right position of the tree
		if val > root.Data {
			root.right = Insert(root.right, val)
		} else { // In case the value is not greater, we assume that is less than the current val, so we insert into the left position
			root.left = Insert(root.left, val)
		}

		// return the current node
		return root
	}
}

// Left takes the node in left side of the binary tree
func (n *Node) Left() *Node {
	return n.left
}

// Right takes the node in right side of the binary tree
func (n *Node) Right() *Node {
	return n.right
}

// Search will look will recursively look for the value inside the tree and return the node containing the val if exists
func (n *Node) Search(val int) *Node {
	// if the current node does not exist, or if it's equal the val we are looking for, we early return the current node
	if n == nil || n.Data == val {
		return n
	}

	// we check if the desired val is greater than the current value, which will recursively call the search into the right node
	if val > n.Data {
		out := fmt.Sprintf("going right of %d", n.Data)
		fmt.Println(out)
		return n.Right().Search(val)
	}

	// if the desired val is not found we assume that is on the left side and go to the left
	out := fmt.Sprintf("going left of %d", n.Data)
	fmt.Println(out)
	return n.Left().Search(val)
}
