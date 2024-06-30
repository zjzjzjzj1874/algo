package middle

// 86. 分隔链表
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
// 示例 1：
// 输入：head = [1,4,3,2,5,2], x = 3
// 输出：[1,2,2,4,3,5]
// 示例 2：
// 输入：head = [2,1], x = 2
// 输出：[1,2]
// 提示：
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：使用4个空指针，大于等于区首尾，小于区首尾，遍历后归属，整合
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		lh, lt     *ListNode // lh:less head;lt less tail
		gteh, gtet *ListNode // gteh:greater or equal head	; gtet:greater or equal tail
	)

	cur := head
	for cur != nil {
		switch {
		case cur.Val < x:
			if lh == nil { // 头尾都为空，首次走这里
				lh = cur
				lt = cur
				cur = cur.Next
				lh.Next = nil
				lt.Next = nil
				continue
			}

			// 不止一个值后：
			lt.Next = cur
			cur = cur.Next
			lt = lt.Next
			lt.Next = nil
		default: // case cur.Val >= x:
			if gteh == nil { // 头尾都为空，首次走这里
				gteh = cur
				gtet = cur
				cur = cur.Next
				gteh.Next = nil
				gtet.Next = nil
				continue
			}

			// 不止一个值后：
			gtet.Next = cur
			cur = cur.Next
			gtet = gtet.Next
			gtet.Next = nil
		}

	}

	//	把四个节点串起来
	if lh == nil { // 小于区没有节点，返回大于等于区
		return gteh
	}
	if gteh == nil { // 大于等于区没有节点，返回小于区
		return lh
	}
	//	大于小于区拼接
	lt.Next = gteh

	return lh
}
