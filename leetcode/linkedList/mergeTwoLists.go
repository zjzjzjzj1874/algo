package linkedList

// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//示例 1：
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//示例 2：
//输入：l1 = [], l2 = []
//输出：[]
//示例 3：
//输入：l1 = [], l2 = [0]
//输出：[0]
//提示：
//两个链表的节点数目范围是 [0, 50]
//-100 <= Node.val <= 100
//l1 和 l2 均按 非递减顺序 排列

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	ans := &ListNode{} // 一个空头节点
//	cur := ans
//
//	for list1 != nil && list2 != nil {
//		if list1.Val < list2.Val {
//			cur.Next = list1
//			list1 = list1.Next
//		} else {
//			cur.Next = list2
//			list2 = list2.Next
//		}
//		cur = cur.Next
//	}
//
//  NOTE!!! 是if    不是for
//	if list1 != nil {
//		cur.Next = list1
//	}
//	if list2 != nil {
//		cur.Next = list2
//	}
//
//	return ans.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}
