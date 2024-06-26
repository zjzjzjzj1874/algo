package tree

import "container/list"

// 145. 二叉树的后序遍历
// 给你二叉树的根节点 root ，返回它节点值的 后序 遍历。
//示例 1：
//输入：root = [1,null,2,3]
//输出：[3,2,1]
//示例 2：
//输入：root = []
//输出：[]
//示例 3：
//输入：root = [1]
//输出：[1]
//提示：
//树中节点数目在范围 [0, 100] 内
//-100 <= Node.val <= 100
//进阶：递归算法很简单，你可以通过迭代算法完成吗？
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	var postorder func(root *TreeNode)

	ans := make([]int, 0, 4)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)
		postorder(root.Right)
		ans = append(ans, root.Val)
	}

	postorder(root)

	return ans
}

func postorderTraversalStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0, 4)
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)

		root = e.Value.(*TreeNode)
		ans = append(ans, root.Val)

		if root.Left != nil {
			stack.PushBack(root.Left)
		}
		if root.Right != nil {
			stack.PushBack(root.Right)
		}
	}

	n := len(ans) // 交换顺序
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}

	return ans
}
