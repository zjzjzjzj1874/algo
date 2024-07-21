package tree

// LCR 053. 二叉搜索树中的中序后继
// 给定一棵二叉搜索树和其中的一个节点 p ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
//
// 节点 p 的后继是值比 p.val 大的节点中键值最小的节点，即按中序遍历的顺序节点 p 的下一个节点。
//
// 示例 1：
//
// 输入：root = [2,1,3], p = 1
// 输出：2
// 解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。
// 示例 2：
//
// 输入：root = [5,3,6,2,4,null,null,1], p = 6
// 输出：null
// 解释：因为给出的节点没有中序后继，所以答案就返回 null 了。
//
// 提示：
//
// 树中节点的数目在范围 [1, 104] 内。
// -105 <= Node.val <= 105
// 树中各节点的值均保证唯一。
//
// 注意：本题与主站 285 题相同： https://leetcode-cn.com/problems/inorder-successor-in-bst/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	nodes := make([]*TreeNode, 0, 4)
	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil { // 空节点或者已经找到结果
			return
		}
		inorder(root.Left)
		nodes = append(nodes, root)
		inorder(root.Right)
	}

	inorder(root)

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Val != p.Val {
			continue
		}

		if i != len(nodes)-1 {
			return nodes[i+1]
		} else {
			return nil
		}
	}

	return nil
}
