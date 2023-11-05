package trees

// Transversal is the representation of a tree Transversal with its root node
type Transversal struct {
	root *Node
}

// NewTransversal Creates the new Transversal from the root
func NewTransversal(root *Node) *Transversal {
	return &Transversal{root: root}
}

// InOrder returns the BST as a slice of integer in-order, which is Left, Root, Right
func (t *Transversal) InOrder() []int {
	output := []int{}

	var inOrder func(node *Node)

	inOrder = func(node *Node) {
		if node == nil {
			return
		}
		inOrder(node.Left())
		output = append(output, node.Data)
		inOrder(node.Right())
	}

	inOrder(t.root)

	return output
}

// PreOrder returns the BST as a slice of integer in pre-order, which is Root, Left, Right
func (t *Transversal) PreOrder() []int {
	output := []int{}

	var preOrder func(node *Node)

	preOrder = func(node *Node) {
		if node == nil {
			return
		}
		output = append(output, node.Data)
		preOrder(node.Left())
		preOrder(node.Right())
	}

	preOrder(t.root)

	return output
}

// PostOrder returns the BST as a slice of integer in post-order, which is Left, Right, Root
func (t *Transversal) PostOrder() []int {
	output := []int{}

	var postOrder func(node *Node)

	postOrder = func(node *Node) {
		if node == nil {
			return
		}
		postOrder(node.Left())
		postOrder(node.Right())
		output = append(output, node.Data)
	}

	postOrder(t.root)

	return output
}

// LevelOrder returns the BST as a slice of integers in level-order
func (t *Transversal) LevelOrder() []int {
	output := []int{}

	// Check if the root is nil, which means the tree is empty
	if t.root == nil {
		return output
	}

	// Initialize a queue with the root node
	queue := []*Node{t.root}

	// Continue looping until the queue is empty
	for len(queue) > 0 {
		// Dequeue the first node from the queue
		node := queue[0]
		queue = queue[1:]

		// Append the data of the node to the output slice
		output = append(output, node.Data)

		// Enqueue the left child if it exists
		if node.Left() != nil {
			queue = append(queue, node.Left())
		}

		// Enqueue the right child if it exists
		if node.Right() != nil {
			queue = append(queue, node.Right())
		}
	}

	return output
}

// ReverseLevelOrder returns the BST as a slice of integers in reverse level-order
func (t *Transversal) ReverseLevelOrder() []int {
	var result [][]int
	output := []int{}

	// Check if the root is nil, which means the tree is empty
	if t.root == nil {
		return output
	}

	// Initialize a queue with the root node
	queue := []*Node{t.root}

	// Continue looping until the queue is empty
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			// Dequeue the first node from the queue
			node := queue[0]
			queue = queue[1:]

			// Add the node's value to the current level slice
			currentLevel = append(currentLevel, node.Data)

			// Enqueue the left child if it exists
			if node.Left() != nil {
				queue = append(queue, node.Left())
			}

			// Enqueue the right child if it exists
			if node.Right() != nil {
				queue = append(queue, node.Right())
			}
		}

		// Prepend the current level to the result slice
		result = append([][]int{currentLevel}, result...)
	}

	// Flatten the result to a single slice
	for _, level := range result {
		output = append(output, level...)
	}

	return output
}
