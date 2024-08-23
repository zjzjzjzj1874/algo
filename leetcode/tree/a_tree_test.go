package tree

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// 2476. 二叉搜索树最近节点查询
func Test_closestNodes(t *testing.T) {
	type args struct {
		root    *TreeNode
		queries []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "closestNodes", args: args{
			root: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 9}},
			queries: []int{3},
		}, want: [][]int{{-1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestNodes(tt.args.root, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 050. 路径总和 III
func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "root", args: args{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{Val: 5,
					Left: &TreeNode{Val: 3,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: -2}},
					Right: &TreeNode{Val: 2,
						Right: &TreeNode{Val: 1}}},
				Right: &TreeNode{Val: -3,
					Right: &TreeNode{Val: 11}},
			},
			targetSum: 8},
			want: 3},
		{name: "root", args: args{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{Val: 5,
					Left: &TreeNode{Val: 3,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: -2}},
					Right: &TreeNode{Val: 2,
						Right: &TreeNode{Val: 1}}},
				Right: &TreeNode{Val: -3,
					Right: &TreeNode{Val: 11}},
			},
			targetSum: 8},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 051. 二叉树中的最大路径和
func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "maxPathSum", args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: -2,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:   -1,
						Right: nil,
						Left:  nil,
					},
					Left: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Right: nil,
					Left:  nil,
				},
			},
			Right: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val:   -2,
					Right: nil,
					Left:  nil,
				},
				Right: nil,
			},
		}}, want: 3},
		{name: "maxPathSum", args: args{root: &TreeNode{
			Val: -10,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   15,
					Right: nil,
					Left:  nil,
				},
				Right: &TreeNode{
					Val:   7,
					Right: nil,
					Left:  nil,
				},
			},
		}}, want: 42},
		{name: "maxPathSum", args: args{root: &TreeNode{
			Val:   -3,
			Left:  nil,
			Right: nil}}, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 054. 把二叉搜索树转换为累加树
func Test_convertBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "convertBST", args: args{root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Right: nil,
					Left:  nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
		}}, want: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val:   10,
					Right: nil,
					Left:  nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 053. 二叉搜索树中的中序后继
func Test_inorderSuccessor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "increaseBST", args: args{root: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
		}}, want: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   7,
					Right: nil,
					Left:  nil,
				},
				Right: nil,
			},
			Right: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderSuccessor(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

// LCR 052. 递增顺序搜索树
func Test_increasingBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "increaseBST", args: args{root: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
		}}, want: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   7,
					Right: nil,
					Left:  nil,
				},
				Right: nil,
			},
			Right: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 94. 二叉树的中序遍历
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

// LCR 047. 二叉树剪枝
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

// LCR 049. 求根节点到叶节点数字之和
func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sumNumber", args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil},
		}}, want: 25},
		{name: "sumNumber", args: args{root: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil}},
			Right: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil},
		}}, want: 1026},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 104. 二叉树的最大深度
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

// 145. 二叉树的后序遍历
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

// 144. 二叉树的前序遍历
func Test_preorderTraversalIterate(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
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
		}, wantAns: []int{1, 2, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := preorderTraversalIterate(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("preorderTraversalIterate() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_findTwoSwapped(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{name: "findTwoSwapped", args: args{nums: []int{1, 2, 7, 4, 5, 6, 3, 8}}, want1: 3, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findTwoSwapped(tt.args.nums)
			if got != tt.want {
				t.Errorf("findTwoSwapped() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findTwoSwapped() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// 337. 打家劫舍 III
func Test_rob(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "rob",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   3,
								Left:  &TreeNode{},
								Right: nil},
							Right: nil},
						Right: nil},
					Right: nil,
				}},
			want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.args.root)
			if got != tt.want {
				t.Errorf("findTwoSwapped() got = %v, want %v", got, tt.want)
			}
		})
	}
}
