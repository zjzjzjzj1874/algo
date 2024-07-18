package tree

import "math"

// 是否是满二叉树
// 解题：bfs，广度优先，找左右树要信息：节点数和高度
func isFBTree(root *TreeNode) bool {
	var process func(root *TreeNode) (int, int)

	// 返回，1.节点数量  2.高度
	process = func(root *TreeNode) (nodeNum int, height int) {
		if root == nil {
			return 0, 0
		}

		ln, lh := process(root.Left)
		rn, rh := process(root.Right)

		hei := max(lh, rh)

		return 1 + ln + rn, hei + 1
	}

	num, hei := process(root)
	return num == int(math.Pow(2, float64(hei)))-1
}
