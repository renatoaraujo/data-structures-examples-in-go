package trees_test

import (
	"reflect"
	"testing"

	"github.com/renatoaraujo/data-structures-examples-in-go/trees"
)

func TestTransversal(t *testing.T) {
	tests := []struct {
		name                  string
		arr                   []int
		wantInOrder           []int
		wantPreOrder          []int
		wantPostOrder         []int
		wantLevelOrder        []int
		wantReverseLevelOrder []int
	}{
		{
			"Empty tree",
			[]int{},
			[]int{},
			[]int{},
			[]int{},
			[]int{},
			[]int{},
		},
		{
			"Single node",
			[]int{1},
			[]int{1},
			[]int{1},
			[]int{1},
			[]int{1},
			[]int{1},
		},
		{
			"Three nodes balanced",
			[]int{2, 1, 3},
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := trees.NewBalancedBSTFromSlice(tt.arr)
			tr := trees.NewTransversal(root)
			if got := tr.InOrder(); !reflect.DeepEqual(got, tt.wantInOrder) {
				t.Errorf("Transversal.InOrder() = %v, want %v", got, tt.wantInOrder)
			}
			if got := tr.PreOrder(); !reflect.DeepEqual(got, tt.wantPreOrder) {
				t.Errorf("Transversal.PreOrder() = %v, want %v", got, tt.wantPreOrder)
			}
			if got := tr.PostOrder(); !reflect.DeepEqual(got, tt.wantPostOrder) {
				t.Errorf("Transversal.PostOrder() = %v, want %v", got, tt.wantPostOrder)
			}
			if got := tr.LevelOrder(); !reflect.DeepEqual(got, tt.wantLevelOrder) {
				t.Errorf("Transversal.LevelOrder() = %v, want %v", got, tt.wantLevelOrder)
			}
			if got := tr.ReverseLevelOrder(); !reflect.DeepEqual(got, tt.wantReverseLevelOrder) {
				t.Errorf("Transversal.ReverseLevelOrder() = %v, want %v", got, tt.wantReverseLevelOrder)
			}
		})
	}
}
