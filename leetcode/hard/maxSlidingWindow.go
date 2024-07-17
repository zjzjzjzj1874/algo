package hard

// 239. 滑动窗口最大值
// 标签：双端队列 滑动窗口
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回 滑动窗口中的最大值 。
//
// 示例 1：
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 示例 2：
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 提示：
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length

// 解题：双端队列
func maxSlidingWindow(nums []int, k int) []int {
	// 基于golang切片的双端队列问题 头 => 尾 ； 大 => 小
	n := len(nums)
	if k > n {
		return []int{}
	}

	ans := make([]int, 0, n-k+1)

	queue := make([]int, 0, k)
	l := 0
	r := 0
	for r < n {
		for r-l < k && r < n {
			// queue队列不为空，且当前元素比队列尾部元素大，需要先将尾部元素弹出
			for len(queue) > 0 && nums[r] > nums[queue[len(queue)-1]] {
				queue = queue[:len(queue)-1]
			}
			queue = append(queue, r)
			r++
		}

		if r-l == k {
			ans = append(ans, nums[queue[0]]) // 最大值为队列头部元素
		}
		if queue[0] < r-k+1 {
			queue = queue[1:] // 移除队列头部元素
		}
		l++
	}

	return ans
}

// 解题：
func maxSlidingWindowWithEasy(nums []int, k int) []int {
	// 基于golang切片的双端队列问题 头 => 尾 ； 大 => 小
	n := len(nums)
	if k > n {
		return []int{}
	}

	ans := make([]int, 0, n-k+1)

	queue := make([]int, 0, k)
	for i, num := range nums {
		// 1.从队尾入队列
		for len(queue) > 0 && num > nums[queue[len(queue)-1]] { // 维护队列的单调性
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i) // 将元素放到队列尾部

		// 2.记录结果
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}

		// 3.从对头出队列
		if i-queue[0] >= k-1 {
			queue = queue[1:] // 移除队列头部元素
		}
	}

	return ans
}
