package tree

import "container/list"

// 94. 二叉树的中序遍历
// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//示例 1：
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//示例 2：
//输入：root = []
//输出：[]
//示例 3：
//输入：root = [1]
//输出：[1]
//提示：
//树中节点数目在范围 [0, 100] 内
//-100 <= Node.val <= 100
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return ans
}

// 解题：迭代
// 中序遍历打印顺序：左头右，使用左树分解整棵树，则压栈顺序是左右左右左右。。。
// 每根子树的整棵树左边界进栈，一次弹出的过程中，打印处理，对弹出节点的右树重复；
func inorderTraversalWithStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0, 4)
	stack := list.New()
	cur := root
	// 1.把所有左树压入栈
	for cur != nil {
		stack.PushBack(cur)
		cur = cur.Left
	}

	// 2. 依次弹出，然后打印
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		cur := e.Value.(*TreeNode)
		ans = append(ans, cur.Val)

		cur = cur.Right
		// 3.对右树重复，先把所有左树放入
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		}
	}

	return ans
}
