package slidingWindow

// 绳子最多覆盖点位
// 给定一个有序数组arr，代表数轴上从左到右有n个点arr[0]、arr[1]...arr[n-1]，给定一个正数L,代表一根长度为L的绳子,求绳子最多能覆盖其中的几个点。
func ropeMaxPoint(nums []int, ropeLen int) int {
	n := len(nums)
	if n == 0 || ropeLen == 0 {
		return n
	}

	start := 0
	ans := 0
	for end := 0; end < n; end++ {
		for nums[end]-nums[start] > ropeLen {
			start++
		}
		ans = max(ans, end-start+1)
	}

	return ans
}
