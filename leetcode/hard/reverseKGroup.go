package hard

// 25. K 个一组翻转链表
// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
// 示例 1：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
// 示例 2：
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]
// 提示：
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	cur := head
	n := 0 // 求链表长度
	for cur != nil {
		n++
		cur = cur.Next
	}

	dummy := &ListNode{Next: head} // 哨兵节点
	var pre *ListNode              //
	p0 := dummy
	cur = p0.Next
	for ; n >= k; n -= k { // 剩余个数大于等于K停止，每次N减少K，因为里面还有个K循环
		for j := 0; j < k; j++ { // 每次翻转K个
			next := cur.Next
			cur.Next = pre
			pre = cur

			cur = next
		}

		// 翻转K和K+1的节点
		next := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = next
	}

	return dummy.Next
}
