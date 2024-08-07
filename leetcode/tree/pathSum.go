package tree

// LCR 050. 路径总和 III
// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
// 示例 1：
//
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。
// 示例 2：
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：3
//
// 提示:
//
// 二叉树的节点个数的范围是 [0,1000]
// -109 <= Node.val <= 109
// -1000 <= targetSum <= 1000
//
// 注意：本题与主站 437 题相同：https://leetcode-cn.com/problems/path-sum-iii/
// 动态规划，找左树和右树要信息
func pathSum(root *TreeNode, targetSum int) (ans int) {
	if root == nil {
		return 0
	}

	var process func(root *TreeNode, rest int) int

	process = func(root *TreeNode, rest int) int {
		if root == nil {
			return 0
		}
		ans := 0
		rest -= root.Val
		if rest == 0 {
			ans++
		}

		ans += process(root.Left, rest)
		ans += process(root.Right, rest)

		return ans
	}
	ans = ans + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum) + process(root, targetSum)
	return
}
