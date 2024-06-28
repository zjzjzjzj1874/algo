package easy

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

func inorderTraversalWithAns(root *TreeNode, ans []int) []int {
	if root == nil {
		return []int{}
	}

	inorderTraversalWithAns(root.Left, ans)
	ans = append(ans, root.Val)
	inorderTraversalWithAns(root.Right, ans)
	return ans
}
