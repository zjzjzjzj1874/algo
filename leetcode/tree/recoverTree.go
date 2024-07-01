package tree

//99. 恢复二叉搜索树
//给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
//示例 1：
//输入：root = [1,3,null,null,2]
//输出：[3,1,null,null,2]
//解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
//示例 2：
//输入：root = [3,1,4,null,null,2]
//输出：[2,1,4,null,null,3]
//解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
//提示：
//树上节点的数目在范围 [2, 1000] 内
//-231 <= Node.val <= 231 - 1
//进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？
//
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解题：先中序遍历，获取有序数组，然后找到出错的两个index的值，最后恢复,会用到多次递归。后续学习Morris 中序遍历
func recoverTree(root *TreeNode) {
	arr := make([]int, 0, 4)

	var inorderTraversal func(tree *TreeNode)
	inorderTraversal = func(tree *TreeNode) {
		if tree == nil {
			return
		}

		inorderTraversal(tree.Left)
		arr = append(arr, tree.Val)
		inorderTraversal(tree.Right)
	}
	inorderTraversal(root)

	x, y := findTwoSwapped(arr)

	recover(root, 2, x, y)
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}

	if root.Val == x || root.Val == y {
		if root.Val == y {
			root.Val = x
		} else {
			root.Val = y
		}

		count--

		if count == 0 {
			return
		}
	}

	recover(root.Left, count, x, y)
	recover(root.Right, count, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	x, y := nums[index1], nums[index2]
	return x, y
}
