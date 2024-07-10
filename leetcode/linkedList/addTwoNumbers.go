package linkedList

// 445. 两数相加 II
// 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
// 示例1：
//
// 输入：l1 = [7,2,4,3], l2 = [5,6,4]
// 输出：[7,8,0,7]
// 示例2：
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[8,0,7]
// 示例3：
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
//
// 提示：
//
// 链表的长度范围为 [1, 100]
// 0 <= node.val <= 9
// 输入数据保证链表代表的数字无前导 0
//
// 进阶：如果输入链表不能翻转该如何解决？

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解题： 栈或数组
func addTwoNumbersWithStack(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	a1 := change2Arr(l1)
	a2 := change2Arr(l2)

	n1 := len(a1)
	n2 := len(a2)
	n := n1
	if n < n2 {
		n = n2
	}
	carry := 0
	dummy := &ListNode{}
	cur := dummy
	for i := 0; i < n; i++ {
		sum := carry

		if n1-i-1 >= 0 {
			sum += a1[n1-i-1]
		}

		if n2-i-1 >= 0 {
			sum += a2[n2-i-1]
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return reverse(dummy.Next)
}

// 解题： 翻转链表
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	p1 := reverseList(l1)
	p2 := reverseList(l2)

	carry := 0
	dummy := &ListNode{}
	cur := dummy
	for p1 != nil && p2 != nil {
		sum := p1.Val + p2.Val + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		sum := p1.Val + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		p1 = p1.Next
	}
	for p2 != nil {
		sum := p2.Val + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		p2 = p2.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var dummy *ListNode

	for head != nil {
		next := head.Next
		head.Next = dummy
		dummy = head
		head = next
	}

	return dummy
}

func change2Arr(head *ListNode) (ans []int) {
	if head == nil {
		return
	}

	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}

	return
}
