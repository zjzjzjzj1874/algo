package linkedList

// 148. 排序链表
// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
// 示例 1：
//
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
// 示例 2：
//
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]
// 示例 3：
//
// 输入：head = []
// 输出：[]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 5 * 104] 内
// -105 <= Node.val <= 105
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：按照上一题的插入排序试试,时间复杂度不满足，记得使用归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}

	sorted, cur := head, head.Next
	for cur != nil {
		if sorted.Val <= cur.Val { // 有序
			sorted = sorted.Next
		} else { // 无序，需要找到最接近的位置插入
			pre := dummy
			for pre.Next.Val <= cur.Val {
				pre = pre.Next
			}

			// 找到了
			sorted.Next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur
		}

		cur = sorted.Next
	}

	return dummy.Next
}

func sortListWithMerge(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	return sort(head, nil)
}

func merge(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	// 两个链表都有值
	var (
		dummy  = &ListNode{}
		cur    = dummy
		t1, t2 = head1, head2
	)
	for t1 != nil && t2 != nil {
		if t1.Val <= t2.Val {
			cur.Next = t1
			t1 = t1.Next
		} else {
			cur.Next = t2
			t2 = t2.Next
		}
		cur = cur.Next
	}

	if t1 != nil {
		cur.Next = t1
	}
	if t2 != nil {
		cur.Next = t2
	}

	return dummy.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == tail || head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next

		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow

	return merge(sort(head, mid), sort(mid, tail))
}
