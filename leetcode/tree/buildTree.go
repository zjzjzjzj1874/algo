package tree

//105. 从前序与中序遍历序列构造二叉树
//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
//示例 1:
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//示例 2:
//
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]
//提示:
//
//1 <= preorder.length <= 3000
//inorder.length == preorder.length
//-3000 <= preorder[i], inorder[i] <= 3000
//preorder 和 inorder 均 无重复 元素
//inorder 均出现在 preorder
//preorder 保证 为二叉树的前序遍历序列
//inorder 保证 为二叉树的中序遍历序列
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解题：递归，前序遍历：头左右；中序遍历：左头右 =》 所以可以preorder知道头结点是preorder[0]; TODO 看懂递归的条件
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	i := 0

	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val { // 在中序中找到头结点，然后拼接，头结点中左边是左树的，右边是右树的
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
