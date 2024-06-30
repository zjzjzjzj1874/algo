package middle

// 61. 旋转链表
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
// 示例 1：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
// 示例 2：
// 输入：head = [0,1,2], k = 4
// 输出：[2,0,1]
// 提示：
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 109
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：先看head的长度，然后闭合形成链表，根据k%n=k，然后移动k步，断开链表
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	cur := head
	n := 0
	var last *ListNode
	for cur != nil {
		n++
		last = cur // 尾节点

		cur = cur.Next
	}
	k = n - (k % n) // 形成环，然后取反
	if k == 0 {
		return head
	}

	last.Next = head // 闭环链表
	for i := 0; i < k; i++ {
		last = last.Next
	}

	// 出环，断链
	ret := last.Next
	last.Next = nil
	return ret
}
