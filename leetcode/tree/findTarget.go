package tree

// LCR 056. 两数之和 IV - 输入二叉搜索树
// 给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。假设二叉搜索树中节点的值均唯一。
//
// 示例 1：
//
// 输入: root = [8,6,10,5,7,9,11], k = 12
// 输出: true
// 解释: 节点 5 和节点 7 之和等于 12
// 示例 2：
//
// 输入: root = [8,6,10,5,7,9,11], k = 22
// 输出: false
// 解释: 不存在两个节点值之和为 22 的节点
//
// 提示：
//
// 二叉树的节点个数的范围是  [1, 104].
// -104 <= Node.val <= 104
// root 为二叉搜索树
// -105 <= k <= 105
//
// 注意：本题与主站 653 题相同： https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

func findTarget(root *TreeNode, k int) bool {
	nums := make([]int, 0, 4)

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		nums = append(nums, root.Val)
		inorder(root.Right)
	}
	inorder(root)

	left := 0
	right := len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == k {
			return true
		}
		if nums[left]+nums[right] > k {
			right--
			continue
		}
		left++
	}

	return false
}
