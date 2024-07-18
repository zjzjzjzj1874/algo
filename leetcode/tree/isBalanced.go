package tree

// LCR 176. 判断是否为平衡二叉树
// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
//
// 示例 1:
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true
// 解释：如下图
//
// 示例 2:
//
// 输入：root = [1,2,2,3,3,null,null,4,4]
// 输出：false
// 解释：如下图
//
// 提示：
//
// 0 <= 树的结点个数 <= 10000
// 注意：本题与主站 110 题相同：https://leetcode-cn.com/problems/balanced-binary-tree/
// 判断平衡二叉树，即找左右树要树的高度，俩高度差不超过1则满足
func isBalancedWithLT(root *TreeNode) bool {
	var process func(root *TreeNode) int

	// 返回值，1，树的高度，2是否平衡
	process = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 遇到-1表示不平衡，提前退出
		lhei := process(node.Left)
		if lhei == -1 {
			return -1
		}
		rhei := process(node.Right)
		if rhei == -1 {
			return -1
		}

		maxhei := max(lhei, rhei)
		if maxhei-lhei > 1 || maxhei-rhei > 1 {
			return -1
		}

		return maxhei + 1
	}

	return !(process(root) == -1)
}
func isBalanced(root *TreeNode) bool {
	var process func(root *TreeNode) (int, bool)

	// 返回值，1，树的高度，2是否平衡
	process = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		lhei, lb := process(node.Left)
		if !lb {
			return 0, lb
		}
		rhei, rb := process(node.Right)
		if !rb {
			return 0, rb
		}

		maxhei := max(lhei, rhei)
		if maxhei-lhei > 1 || maxhei-rhei > 1 {
			return maxhei, false
		}

		return maxhei + 1, true
	}

	_, isbalance := process(root)
	return isbalance
}
