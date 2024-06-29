package middle

// 142. 环形链表 II
// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//不允许修改 链表。
//示例 1：
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//示例 2：
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//示例 3：
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
//提示：
//链表中节点的数目范围在范围 [0, 104] 内
//-105 <= Node.val <= 105
//pos 的值为 -1 或者链表中的一个有效索引
//进阶：你是否可以使用 O(1) 空间解决此题？
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解题：使用hash表存储，遇到第一个重复的则是第一个节点，否则不是环形链表
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	listMap := make(map[*ListNode]int)
	idx := 0
	for head != nil {
		if _, ok := listMap[head]; ok {
			return head
		} else {
			listMap[head] = idx
			head = head.Next
		}
	}

	return nil
}

// 空间复杂度O(1)：快慢指针
func detectCycleO1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil { // 无环，不相遇，退出
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast { // 有环，相遇
			tmp := head

			for tmp != slow {
				tmp = tmp.Next
				slow = slow.Next
			}
			return tmp
		}
	}

	return nil
}
