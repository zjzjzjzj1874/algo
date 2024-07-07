package linkedList

// LCR 136. 删除链表的节点
// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
// 返回删除后的链表的头节点。
//
// 示例 1:
//
// 输入: head = [4,5,1,9], val = 5
// 输出: [4,1,9]
// 解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
// 示例 2:
//
// 输入: head = [4,5,1,9], val = 1
// 输出: [4,5,9]
// 解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
//
// 说明：
//
// 题目保证链表中节点的值互不相同
// 若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：链表删除
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	cur := dummy

	for cur.Next != nil {
		if cur.Next.Val == val {
			// 当前节点是需要删除的节点
			cur.Next = cur.Next.Next
			return dummy.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}