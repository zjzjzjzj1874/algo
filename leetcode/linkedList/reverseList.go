package linkedList

// 206. 反转链表
// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]
// 示例 3：
// 输入：head = []
// 输出：[]
// 提示：
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

// 解题：需要使用一个中间变量缓存next的指针
func reverseList(head *ListNode) *ListNode {
	var dummy *ListNode

	cur := head
	for cur != nil {
		next := cur.Next

		cur.Next = dummy
		dummy = cur

		cur = next
	}

	return dummy
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	var ans *ListNode
//	cur := head
//
//	for cur != nil {
//		next := cur.Next
//		cur.Next = ans
//
//		ans = cur
//		cur = next
//	}
//
//	return ans
//}
