package tree

import "container/list"

// 144. 二叉树的前序遍历
// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//示例 1：
//输入：root = [1,null,2,3]
//输出：[1,2,3]
//示例 2：
//输入：root = []
//输出：[]
//示例 3：
//输入：root = [1]
//输出：[1]
//示例 4：
//输入：root = [1,2]
//输出：[1,2]
//示例 5：
//输入：root = [1,null,2]
//输出：[1,2]
//提示：
//树中节点数目在范围 [0, 100] 内
//-100 <= Node.val <= 100
//进阶：递归算法很简单，你可以通过迭代算法完成吗？
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解题：递归序遍历
func preorderTraversal(root *TreeNode) []int {
	var preorder func(root *TreeNode)

	ans := make([]int, 0, 4)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		ans = append(ans, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)

	return ans
}

// 解题：迭代遍历,递归会帮助我们自己压栈，如果使用迭代，则需要自己压栈
// 先序遍历：头左右，则压栈顺序：右左头
// 步骤 1.把根节点放入栈中；2.从栈中弹出一个节点cur；3.处理(打印)弹出的节点；4.先右后左把子节点(如果有)放入栈中；5.重复2-4的步骤；
func preorderTraversalIterate(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	stack := list.New()
	stack.PushBack(root) // 压入双链表尾部
	cur := root
	for stack.Len() > 0 {
		head := stack.Back()
		stack.Remove(head)
		ans = append(ans, head.Value.(*TreeNode).Val)

		cur = head.Value.(*TreeNode) // cur不用判空，因为写入就保证了`!=nil`
		if cur.Right != nil {
			stack.PushBack(cur.Right) // 压入双链表尾部
		}
		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}
	}

	return ans
}

func preorderTraversalArr(root *TreeNode) (ans []int) {
	stack := make([]*TreeNode, 0, 4)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			ans = append(ans, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return ans
}
