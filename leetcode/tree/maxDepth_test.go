package tree

import (
	"fmt"
	"math"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxDepth", args: args{root: &TreeNode{
			Val: 3, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{
				Val: 20, Right: &TreeNode{
					Val: 7, Left: nil, Right: nil}, Left: &TreeNode{
					Val: 15, Left: nil, Right: nil},
			},
		}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("MathPow", func(t *testing.T) {
		fmt.Println(math.Pow(2, 4))
	})
}
