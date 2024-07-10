package linkedList

// 725. 分隔链表
// 给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
//
// 每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
//
// 这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。
//
// 返回一个由上述 k 部分组成的数组。
//
// 示例 1：
//
// 输入：head = [1,2,3], k = 5
// 输出：[[1],[2],[3],[],[]]
// 解释：
// 第一个元素 output[0] 为 output[0].val = 1 ，output[0].next = null 。
// 最后一个元素 output[4] 为 null ，但它作为 ListNode 的字符串表示是 [] 。
// 示例 2：
//
// 输入：head = [1,2,3,4,5,6,7,8,9,10], k = 3
// 输出：[[1,2,3,4],[5,6,7],[8,9,10]]
// 解释：
// 输入被分成了几个连续的部分，并且每部分的长度相差不超过 1 。
// 0 <= Node.val <= 1000
// 1 <= k <= 50
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) (ans []*ListNode) {
	if k == 1 {
		return []*ListNode{head}
	}
	n := 0 // 链表长度
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	num := n / k
	left := n % k // 剩余数量

	cur := head
	for cur != nil {
		pageSize := num // 本次截取的长度
		if left > 0 {
			pageSize++
			left--
		}
		var thisTime = &ListNode{
			Next: cur,
		}
		tmp := thisTime
		for i := 0; i < pageSize; i++ {
			tmp.Next = cur
			tmp = tmp.Next
			cur = cur.Next
		}
		tmp.Next = nil
		ans = append(ans, thisTime.Next)
	}

	if len(ans) < k {
		la := len(ans)
		for i := 0; i < k-la; i++ {
			ans = append(ans, nil)
		}
	}

	return ans
}

// 解题：优化
func splitListToPartsV1(head *ListNode, k int) (ans []*ListNode) {
	if k == 1 {
		return []*ListNode{head}
	}
	n := 0 // 链表长度
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	num := n / k
	left := n % k // 剩余数量

	ans = make([]*ListNode, k)
	for i, cur := 0, head; i < k && cur != nil; i++ {
		pageSize := num // 本次截取的长度
		if left > 0 {
			pageSize++
			left--
		}
		var thisTime = &ListNode{
			Next: cur,
		}
		tmp := thisTime
		for j := 0; j < pageSize; j++ {
			tmp.Next = cur
			tmp = tmp.Next
			cur = cur.Next
		}
		tmp.Next = nil
		ans[i] = thisTime.Next
	}

	return ans
}
