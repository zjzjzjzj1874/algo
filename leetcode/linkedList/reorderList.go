package linkedList

// 143. 重排链表
// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1：
//
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]
// 示例 2：
//
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]
//
// 提示：
//
// 链表的长度范围为 [1, 5 * 104]
// 1 <= node.val <= 1000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：
func reorderList(head *ListNode) {
	nodes := make([]*ListNode, 0, 8)

	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}

	i := 0
	j := len(nodes) - 1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}

		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
