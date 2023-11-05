package trees_test

import (
	"reflect"
	"testing"

	"github.com/renatoaraujo/data-structures-examples-in-go/trees"
)

type bstTestCase struct {
	name  string
	input []int
	want  []int
}

var bstTestCases = []bstTestCase{
	{
		name:  "Empty slice",
		input: []int{},
		want:  nil,
	},
	{
		name:  "One element",
		input: []int{1},
		want:  []int{1},
	},
	{
		name:  "Multiple elements",
		input: []int{1, 2, 3, 4, 5, 6, 7},
		want:  []int{1, 2, 3, 4, 5, 6, 7},
	},
	{
		name:  "Duplicates",
		input: []int{1, 1, 2, 2, 3, 3},
		want:  []int{1, 1, 2, 2, 3, 3},
	},
	{
		name:  "Unsorted",
		input: []int{4, 2, 6, 1, 3, 5, 7},
		want:  []int{1, 2, 3, 4, 5, 6, 7},
	},
	{
		name:  "Negative elements",
		input: []int{-7, -6, -5, -4, -3, -2, -1},
		want:  []int{-7, -6, -5, -4, -3, -2, -1},
	},
	{
		name:  "Complex",
		input: []int{10, -10, 20, 0, 30, -20, -30},
		want:  []int{-30, -20, -10, 0, 10, 20, 30},
	},
}

func TestNewBalancedBSTFromSlice(t *testing.T) {
	for _, tc := range bstTestCases {
		t.Run(tc.name, func(t *testing.T) {
			transversal := trees.NewTransversal(trees.NewBalancedBSTFromSlice(tc.input))
			result := transversal.InOrder()
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("%s: NewBalancedBSTFromSlice(%v) = %v, want %v", tc.name, tc.input, result, tc.want)
			}
		})
	}
}
