package main

import (
	"fmt"
	"github.com/renatoaraujo/data-structures-examples-in-go/trees"
)

func main() {
	root := &trees.Node{Data: 20}
	trees.Insert(root, 10)
	trees.Insert(root, 25)
	trees.Insert(root, 5)
	trees.Insert(root, 13)
	trees.Insert(root, 7)
	trees.Insert(root, 6)
	trees.Insert(root, 8)
	trees.Insert(root, 23)
	trees.Insert(root, 24)

	//arr := []int{1, 2, 234, 142, 1, 5, 2352, 46, 2346, 3467, 34, 6, 36, 34, 63, 46, 347, 34, 25, 2134, 12}
	//root := trees.NewBalancedBSTFromSlice(arr)

	//val := 2134
	//found := root.Search(val)
	//if found == nil {
	//	out := fmt.Sprintf("values does not exist in the tree: %d", val)
	//	fmt.Println(out)
	//	out = fmt.Sprintf("inserting: %d", val)
	//	fmt.Println(out)
	//	trees.Insert(root, val)
	//
	//	found = root.Search(val)
	//	out = fmt.Sprintf("value found after inserting: %d", val)
	//	fmt.Println(out)
	//} else {
	//	out := fmt.Sprintf("already exists in the tree: %d", val)
	//	fmt.Println(out)
	//}

	transversal := trees.NewTransversal(root)
	inOrder := transversal.InOrder()
	fmt.Println(inOrder)

	preOrder := transversal.PreOrder()
	fmt.Println(preOrder)

	postOrder := transversal.PostOrder()
	fmt.Println(postOrder)

	levelOrder := transversal.LevelOrder()
	fmt.Println(levelOrder)

	reverseLevelOrder := transversal.ReverseLevelOrder()
	fmt.Println(reverseLevelOrder)
}
