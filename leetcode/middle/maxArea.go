package middle

// 11. 盛最多水的容器
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
// 示例 1：
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 示例 2：
// 输入：height = [1,1]
// 输出：1
// 提示：
// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104

// 解题：双指针，左右两个指针，碰到大的就固定，让小的左移或者右移，相等的情况下可随意移动
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0

	for left <= right {
		step := right - left
		min := height[left]
		if height[left] > height[right] {
			min = height[right]
			right--
		} else {
			left++
		}
		if max < step*min {
			max = step * min
		}
	}

	return max
}
