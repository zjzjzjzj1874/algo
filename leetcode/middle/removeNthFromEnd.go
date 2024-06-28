package middle

// 19. 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 示例 1：
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]
// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]
// 提示：
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

//进阶：你能尝试使用一趟扫描实现吗？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：大部分链表都需要构造一个哑结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ori := head

	l := 0
	for head != nil { // 求出其长度
		l++
		head = head.Next
	}

	dummy := &ListNode{Next: ori} // 构造哑结点
	cur := dummy
	for i := 0; i < l-n; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return dummy.Next
}

// 使用栈来做
func removeNthFromEndWithStack(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // 构造哑结点
	cur := dummy
	stacks := make([]*ListNode, 0, 4)
	for cur != nil {
		stacks = append(stacks, cur)
		cur = cur.Next
	}

	remove := stacks[len(stacks)-n-1]
	remove.Next = remove.Next.Next

	return dummy.Next
}
