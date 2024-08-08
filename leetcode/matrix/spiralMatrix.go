package matrix

// 2326. 螺旋矩阵 IV
// 给你两个整数：m 和 n ，表示矩阵的维数。
//
// 另给你一个整数链表的头节点 head 。
//
// 请你生成一个大小为 m x n 的螺旋矩阵，矩阵包含链表中的所有整数。链表中的整数从矩阵 左上角 开始、顺时针 按 螺旋 顺序填充。如果还存在剩余的空格，则用 -1 填充。
//
// 返回生成的矩阵。
//
// 示例 1：
//
// 输入：m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
// 输出：[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
// 解释：上图展示了链表中的整数在矩阵中是如何排布的。
// 注意，矩阵中剩下的空格用 -1 填充。
// 示例 2：
//
// 输入：m = 1, n = 4, head = [0,1,2]
// 输出：[[0,1,2,-1]]
// 解释：上图展示了链表中的整数在矩阵中是如何从左到右排布的。
// 注意，矩阵中剩下的空格用 -1 填充。
//
// 提示：
//
// 1 <= m, n <= 105
// 1 <= m * n <= 105
// 链表中节点数目在范围 [1, m * n] 内
// 0 <= Node.val <= 1000
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func spiralMatrix(m int, n int, head *ListNode) (ans [][]int) {
	if head == nil {
		return
	}

	left := 0
	right := n - 1
	top := 0
	bottom := m - 1

	count := m * n
	ans = make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	cur := head
	for top <= bottom && left <= right {
		for i := top; i <= right; i++ {
			if cur != nil {
				ans[top][i] = cur.Val
				cur = cur.Next
			} else {
				ans[top][i] = -1
			}
			count--
		}
		top++
		if count == 0 {
			return
		}

		for i := top; i <= bottom; i++ {
			if cur != nil {
				ans[i][right] = cur.Val
				cur = cur.Next
			} else {
				ans[i][right] = -1
			}
			count--
		}
		right--
		if count == 0 {
			return
		}

		for i := right; i >= left; i-- {
			if cur != nil {
				ans[bottom][i] = cur.Val
				cur = cur.Next
			} else {
				ans[bottom][i] = -1
			}
			count--
		}
		bottom--
		if count == 0 {
			return
		}

		for i := bottom; i >= top; i-- {
			if cur != nil {
				ans[i][left] = cur.Val
				cur = cur.Next
			} else {
				ans[i][left] = -1
			}
			count--
		}
		left++
		if count == 0 {
			return
		}
	}
	return
}
