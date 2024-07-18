package tree

// 958. 二叉树的完全性检验
// 给你一棵二叉树的根节点 root ，请你判断这棵树是否是一棵 完全二叉树 。
//
// 在一棵 完全二叉树 中，除了最后一层外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。最后一层（第 h 层）中可以包含 1 到 2h 个节点。
//
// 示例 1：
//
// 输入：root = [1,2,3,4,5,6]
// 输出：true
// 解释：最后一层前的每一层都是满的（即，节点值为 {1} 和 {2,3} 的两层），且最后一层中的所有节点（{4,5,6}）尽可能靠左。
// 示例 2：
//
// 输入：root = [1,2,3,4,5,null,7]
// 输出：false
// 解释：值为 7 的节点不满足条件「节点尽可能靠左」。
//
// 提示：
//
// 树中节点数目在范围 [1, 100] 内
// 1 <= Node.val <= 1000

// 解题：广度优先遍历 ==> 二叉树层序遍历，遍历节点之前不能遇到空节点，否则表示不完全
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nodes := make([]*TreeNode, 1, 4) // 构造栈
	nodes[0] = root
	empty := false // 是否遇到空节点
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:] // 取出第一个节点

		if node == nil {
			// 遇到空节点了
			empty = true
		} else {
			// 遍历非空节点，如果遇到空节点，表示之前出现过空节点，不满足完全二叉树的定义，直接return false
			if empty {
				return false
			}

			nodes = append(nodes, node.Left)
			nodes = append(nodes, node.Right)
		}
	}

	return true
}

// 解题：深度优先遍历dfs
func isCompleteTreeWithDFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(root *TreeNode, curNodeIdx, nodeNum int) bool
	dfs = func(root *TreeNode, curNodeIdx, nodeNum int) bool {
		if root == nil {
			return true
		}

		// 节点编号 > 节点数量，表示有空节点
		if curNodeIdx > nodeNum {
			return false
		}

		return dfs(root.Left, curNodeIdx*2, nodeNum) && dfs(root.Right, curNodeIdx*2+1, nodeNum)
	}

	return dfs(root, 1, getNodes(root))
}

// 获取节点数量
func getNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ln := getNodes(root.Left)
	rn := getNodes(root.Right)

	return 1 + ln + rn
}
