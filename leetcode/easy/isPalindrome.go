package easy

// 234. 回文链表
// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
// 示例 1：
// 输入：head = [1,2,2,1]
// 输出：true
// 示例 2：
// 输入：head = [1,2]
// 输出：false
// 提示：
// 链表中节点数目在范围[1, 105] 内
// 0 <= Node.val <= 9
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：使用数组
func isPalindrome(head *ListNode) bool {
	cur := head
	n := 0

	lists := make([]*ListNode, 0, 4)
	for cur != nil {
		lists = append(lists, cur)
		cur = cur.Next
		n++
	}

	for i := 0; i < n/2; i++ {
		if lists[n-1-i].Val != head.Val {
			return false
		}
		head = head.Next
	}

	return true
}

// TODO O(1)复杂度的算法：使用快慢指针，1.找到一半长度的链表；2.反转后半部分的链表；3.比较前后段部分是否相同，返回结果

// 解题：进阶 =》快慢指针
func isPalindromeFastSlow(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 这里慢指针走到一半了，快指针走完了

	return true
}
