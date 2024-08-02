package middle

// 求小和问题
// 在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。求一个数组的小和。
// 解题：反向：从左往右，每个数比右边小的有多少，乘上小数的值，则为这个数，然后对这些数求和。
func smallSum(nums []int) int {
	return process(nums, 0, len(nums)-1)
}

func process(nums []int, l, r int) int {
	if l == r {
		return 0
	}

	mid := l + (r-l)/2

	return process(nums, l, mid) + process(nums, mid+1, r) + mergeSmallSum(nums, l, mid, r)
}

func mergeSmallSum(nums []int, l, m, r int) int {
	helper := make([]int, r-l+1)

	i := 0 // helper数组的下标
	p1 := l
	p2 := m + 1
	ans := 0 // 返回的小和

	for p1 <= m && p2 <= r {
		if nums[p1] < nums[p2] { // 一定要严格小于，
			ans += (r - p2 + 1) * nums[p1] // 计算左边比右边小的值 = 数*值
			helper[i] = nums[p1]
			p1++
		} else {
			helper[i] = nums[p2]
			p2++
		}
		i++
	}

	for p1 <= m {
		helper[i] = nums[p1]
		i++
		p1++
	}

	for p2 <= r {
		helper[i] = nums[p2]
		p2++
		i++
	}

	for i := 0; i < len(helper); i++ {
		nums[i+l] = helper[i]
	}

	return ans
}
