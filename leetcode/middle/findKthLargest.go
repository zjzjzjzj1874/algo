package middle

// 215. 数组中的第K个最大元素
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
//
// 提示：
//
// 1 <= k <= nums.length <= 105
// -104 <= nums[i] <= 104

// 解题：堆排
func findKthLargestWithHeapSort(nums []int, k int) int {
	n := len(nums)
	if n == 0 || k > n {
		return -1
	}

	// 构建大根堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	for i := n - 1; i >= n-k+1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	return nums[0]
}

func heapify(nums []int, n, i int) {
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && nums[left] > nums[largest] {
			largest = left
		}
		if right < n && nums[right] > nums[largest] {
			largest = right
		}
		// 上面是i和左右孩子pk最大值

		if largest == i { // pk出来i最大，啥都不做
			break
		}

		nums[i], nums[largest] = nums[largest], nums[i] // i和最大值交换
		i = largest
	}
}

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if n == 0 || k > n {
		return -1
	}

	quickHelperV2(nums, 0, n-1)

	return nums[n-k]
}

// partitionV2返回一个数组，等于区起始位置
func quickHelperV2(nums []int, l, r int) {
	if l < r {
		idx := partitionV2(nums, l, r)
		quickHelperV2(nums, l, idx[0]-1)
		quickHelperV2(nums, idx[1]+1, r)
	}
}

// 荷兰国旗，小于 等于 大于区
func partitionV2(nums []int, l, r int) []int {
	less := l - 1   // <区右边界
	more := r + 1   // 大于区左边界
	base := nums[r] // 基准值

	for l < more {
		if nums[l] < base {
			nums[l], nums[less+1] = nums[less+1], nums[l]
			less++
			l++
		} else if nums[l] == base {
			// 相等，l++
			l++
		} else {
			nums[more-1], nums[l] = nums[l], nums[more-1]
			more--
		}
	}

	return []int{less + 1, more - 1}
}
