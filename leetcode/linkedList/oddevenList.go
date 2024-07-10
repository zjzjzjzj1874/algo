package linkedList

// 328. 奇偶链表
// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
//
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
//
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
//
// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
//
// 示例 1:
//
// 输入: head = [1,2,3,4,5]
// 输出: [1,3,5,2,4]
// 示例 2:
//
// 输入: head = [2,1,3,5,6,4,7]
// 输出: [2,3,6,7,1,5,4]
//
// 提示:
//
// n ==  链表中的节点数
// 0 <= n <= 104
// -106 <= Node.val <= 106
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	oddDummy := &ListNode{}
	odd := oddDummy
	evenDummy := &ListNode{}
	even := evenDummy
	n := 1

	cur := head
	for cur != nil {
		next := cur.Next
		switch n % 2 {
		case 1:
			odd.Next = cur
			odd.Next.Next = nil
			odd = odd.Next
		case 0:
			even.Next = cur
			even.Next.Next = nil
			even = even.Next
		default:
			return nil
		}
		cur = next
		n++
	}

	odd.Next = evenDummy.Next

	return oddDummy.Next
}

func oddEvenListV1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	evenHead := head.Next
	odd := head
	even := evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
