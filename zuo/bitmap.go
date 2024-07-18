package zuo

// 求0-40亿之间没有出现过的数字,假设数组中的元素不重复
// 解题1：异或
// 思路：
//
//	1.计算0到N之间所有数的异或值，记为 totalXor。
//	2.计算数组中所有数的异或值，记为 arrXor。
// ·3.缺失的数就是 totalXor ^ arrXor。

func findMissingNumberXor(nums []int) int {
	n := len(nums)

	totalXor := 0
	numXor := 0

	for i := 0; i <= n; i++ {
		totalXor ^= i
	}
	for i := 0; i < n; i++ {
		numXor ^= nums[i]
	}

	return totalXor ^ numXor
}

// 解题2: 如果有序，可以使用二分法
func findMissingNumberDichotomy(nums []int) int {
	n := len(nums)
	start := 0
	end := n - 1

	for start <= end {
		mid := start + (end-start)/2

		// 如果中间位置的值等于索引值，缺失的数在右侧
		if nums[mid] == mid {
			start = mid + 1
		} else {
			// 否则，缺失的数在左侧
			end = mid - 1
		}
	}

	return start
}

// 解题3: 求和法,先计算0-n的和，再计算数组的和，两者之差即为缺失的数
func findMissingNumberSum(nums []int) int {
	n := len(nums)

	// 计算0到N的和
	sum := uint64(n) * uint64(n+1) / 2

	for i := 0; i < n; i++ {
		sum -= uint64(nums[i])
	}

	return int(sum)
}
