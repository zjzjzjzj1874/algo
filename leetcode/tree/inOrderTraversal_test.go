package tree

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "inorder", args: args{root: &TreeNode{
			Val: 1, Left: nil, Right: &TreeNode{
				Val: 2, Right: nil, Left: &TreeNode{
					Val: 3, Left: nil, Right: nil},
			},
		}}, want: []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalWithStack(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "pruneTree", args: args{root: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil},
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil},
			},
		}}, want: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Left: nil,
			},
		}},
		{name: "pruneTree", args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil},
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
		}}, want: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Left: nil,
			},
		}},
		{name: "pruneTree", args: args{root: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil},
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		}}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTreeWithLT(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				//if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
