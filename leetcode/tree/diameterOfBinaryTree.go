package tree

// 543. 二叉树的直径 - 简单
// 给你一棵二叉树的根节点，返回该树的 直径 。
//
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
//
// 两节点之间路径的 长度 由它们之间边数表示。
//
// 示例 1：
//
// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
// 示例 2：
//
// 输入：root = [1,2]
// 输出：1
//
// 提示：
//
// 树中节点数目在范围 [1, 104] 内
// -100 <= Node.val <= 100
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTreeWithLT(root *TreeNode) int {
	ans := 0
	var process func(root *TreeNode) int

	process = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		lmax := process(node.Left) + 1
		rmax := process(node.Right) + 1

		ans = max(ans, lmax+rmax)
		return max(lmax, rmax)
	}

	process(root)

	return ans
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans, _ := process(root)

	return ans - 1
}

// 递归处理，p1：左树的最大距离  p2：右树的最大距离 p3：左右两树的最大距离
func process(root *TreeNode) (maxDistance, height int) {
	if root == nil {
		return
	}

	p1, lheight := process(root.Left)
	p2, rheight := process(root.Right)

	p3 := lheight + rheight + 1

	maxDistance = max(p1, max(p2, p3))
	height = max(lheight, rheight) + 1

	return
}
