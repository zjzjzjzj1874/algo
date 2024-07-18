package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris 二叉树的Morris遍历：1977年Morris发明的,遍历的流程：
// 1.如果cur没有左孩子，cur向右移动(cur = cur.Right)
// 2.如果cur有左孩子，找到左子树的最右节点(mostRight)
//
//	2.1 如果mostRight的右指针指向nil，让其指向cur，然后cur向左移动(cur = cur.Left)
//	2.2 如果mostRight的右指针指向cur，让其指向nil，然后cur向右移动(cur = cur.Right)
func Morris(root *TreeNode) {
	var cur, mostRight *TreeNode = root, nil
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil { // 有左孩子
			// 2. cur有左孩子，找左孩子的最右节点；
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			// 2.1 mostRight的右指针指向nil，让其指向cur，cur左移动
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else { // 2.2 mostRight的右指针不为空，表示其指向cur，让其指向nil，然后cur向右移动
				mostRight.Right = nil
			}
		}

		cur = cur.Right // 1. cur没有左孩子，cur向右移动
	}
}
