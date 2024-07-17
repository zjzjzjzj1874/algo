package tree

// 337. 打家劫舍 III
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
// 示例 1:
//
// 输入: root = [3,2,3,null,3,null,1]
// 输出: 7
// 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
// 示例 2:
//
// 输入: root = [3,4,5,1,3,null,1]
// 输出: 9
// 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
//
// 提示：
//
// 树的节点数在 [1, 104] 范围内
// 0 <= Node.val <= 104

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解题：假设根节点是x，
// 则：偷x的情况下的最大值：x+左树的最大值+右树的最大值
// 不偷x：左树+左树最大值+右树最大值
func rob(root *TreeNode) int {
	steal, unSteal := robProcess(root)
	if steal > unSteal {
		return steal
	}
	return unSteal
}

// 返回两个参数，偷的最大值，不偷的最大值
func robProcess(root *TreeNode) (steal int, unSteal int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, 0
	}

	ls, lus := 0, 0
	rs, rus := 0, 0
	if root.Left != nil {
		ls, lus = robProcess(root.Left)
	}
	if root.Right != nil {
		rs, rus = robProcess(root.Right)
	}

	steal = root.Val + lus + rus // 偷这家，那么后面两家就不能偷

	unSteal = max(ls, lus) + max(rs, rus) // 不偷这家，那么不偷的最大值等于倆子树的最大值之和
	return
}
