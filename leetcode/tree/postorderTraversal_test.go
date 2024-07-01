package tree

import (
	"reflect"
	"testing"
)

func Test_postorderTraversalStack(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "preorder", args: args{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		}, want: []int{4, 2, 5, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversalStack(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversalStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
