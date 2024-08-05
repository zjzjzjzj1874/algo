package linkedList

//82. 删除排序链表中的重复元素 II
//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//示例 1：
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//示例 2：
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//提示：
//链表中节点数目在范围 [0, 300] 内
//-100 <= Node.val <= 100
//题目数据保证链表已经按升序 排列
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head} // 哨兵节点，因为头结点也可能重复，被删掉
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 值相等了，找最后相等的
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next // 不相等时移动指针
		}
	}

	return dummy.Next
}

// 83. 删除排序链表中的重复元素
// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//
// 示例 1：
//
// 输入：head = [1,1,2]
// 输出：[1,2]
// 示例 2：
//
// 输入：head = [1,1,2,3,3]
// 输出：[1,2,3]
//
// 提示：
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
func deleteDuplicatesI(head *ListNode) *ListNode {
	eleMap := make(map[int]struct{})

	dummy := &ListNode{}
	nh := dummy
	for cur := head; cur != nil; cur = cur.Next {
		_, ok := eleMap[cur.Val]
		if ok { // 元素重复，需要去重
			continue
		} else {
			nh.Next = cur
			nh = nh.Next
			eleMap[cur.Val] = struct{}{}
		}
	}
	nh.Next = nil

	return dummy.Next
}

func deleteDuplicatesIO1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head} // 哨兵节点，因为头结点也可能重复，被删掉
	cur := dummy

	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next // 当前节点和下一个节点值相等，不移动指针，让下一个指针指向下下个就可以了
		} else {
			cur = cur.Next // 不相等时移动指针
		}
	}

	return dummy.Next
}
