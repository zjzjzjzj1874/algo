package tree

// LCR 047. 二叉树剪枝
// 给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
//
// 节点 node 的子树为 node 本身，以及所有 node 的后代。
//
// 示例 1:
//
// 输入: [1,null,0,0,1]
// 输出: [1,null,0,null,1]
// 解释:
// 只有红色节点满足条件“所有不包含 1 的子树”。
// 右图为返回的答案。
//
// 示例 2:
//
// 输入: [1,0,1,0,0,0,1]
// 输出: [1,null,1,null,1]
// 解释:
//
// 示例 3:
//
// 输入: [1,1,0,1,1,0,1,0]
// 输出: [1,1,0,1,1,null,1]
// 解释:
//
// 提示:
//
// 二叉树的节点个数的范围是 [1,200]
// 二叉树节点的值只会是 0 或 1
//
// 注意：本题与主站 814 题相同：https://leetcode-cn.com/problems/binary-tree-pruning/

func pruneTreeWithLT(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTreeWithLT(root.Left)
	root.Right = pruneTreeWithLT(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var preOrder func(node *TreeNode) bool
	preOrder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if node.Val == 0 && node.Left == nil && node.Right == nil {
			return true
		}
		lPrune := false
		rPrune := false
		if node.Left != nil {
			lPrune = preOrder(node.Left)
			if lPrune {
				node.Left = nil
			}
		} else {
			lPrune = true
		}
		if node.Right != nil {
			rPrune = preOrder(node.Right)
			if rPrune {
				node.Right = nil
			}
		} else {
			rPrune = true
		}
		if lPrune && rPrune && node.Val == 0 {
			return true
		}

		return false
	}

	prune := preOrder(root)
	if prune {
		return nil
	}

	return root
}
