package tree

// 101. 对称二叉树
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
//示例 1：
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//示例 2：
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//提示：
//树中节点数目在范围 [1, 1000] 内
//-100 <= Node.val <= 100
//进阶：你可以运用递归和迭代两种方法解决这个问题吗？
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解题：递归，左右两棵树同时相等
func isSymmetric(root *TreeNode) bool {
	return equal(root, root)
}

func equal(left, right *TreeNode) bool {
	if left == nil && right == nil { // 同时为空，相等
		return true
	}

	// 有一棵子树不为空，另一个为空，不相等
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && equal(left.Left, right.Right) && equal(left.Right, right.Left)
}
