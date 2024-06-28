package middle

// 24. 两两交换链表中的节点
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// 示例 2：
// 输入：head = []
// 输出：[]
// 示例 3：
// 输入：head = [1]
// 输出：[1]
// 提示：
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 解题：迭代
func swapPairsIterate(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	tmp := dummy
	for tmp.Next != nil && tmp.Next.Next != nil {
		l1 := tmp.Next
		l2 := tmp.Next.Next

		// 交换Next位置
		tmp.Next = l2
		l1.Next = l2.Next
		l2.Next = l1

		tmp = l1
	}

	return dummy.Next
}
