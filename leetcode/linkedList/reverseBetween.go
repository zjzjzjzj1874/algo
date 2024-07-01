package linkedList

//92. 反转链表 II
//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
//示例 1：
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//示例 2：
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//提示：
//
//链表中节点数目为 n
//1 <= n <= 500
//-500 <= Node.val <= 500
//1 <= left <= right <= n
//
//
//进阶： 你可以使用一趟扫描完成反转吗？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head} // 哨兵节点
	p0 := dummy

	for i := 0; i < left-1; i++ {
		p0 = p0.Next // 这里先来到翻转的前一个节点
	}

	cur := p0.Next // 来到正式需要翻转的节点
	var pre *ListNode
	for i := 0; i <= right-left; i++ { // 将left,right的节点用来翻转，作为一个临时变量pre
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	} // 翻转小循环

	//拼接三部分的值
	p0.Next.Next = cur // 这两步顺序不能乱，先把剩余的部分拼接，再拼接pre(已经翻转过的)
	p0.Next = pre
	return dummy.Next
}
