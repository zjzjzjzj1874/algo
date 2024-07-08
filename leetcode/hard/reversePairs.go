package hard

// LCR 170. 交易逆序对的总数
// 在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。
//
// 示例 1:
//
// 输入：record = [9, 7, 5, 4, 6]
// 输出：8
// 解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。
//
// 限制：
//
// 0 <= record.length <= 50000

// 解题：归并排序，实质即：求前一个数比后面数大
func reversePairs(record []int) int {
	return process(record, 0, len(record)-1)
}

func process(nums []int, l, r int) int {
	if l >= r {
		return 0
	}

	mid := l + (r-l)/2

	return process(nums, l, mid) + process(nums, mid+1, r) + mergePairs(nums, l, mid, r)
}

func mergePairs(nums []int, l, m, r int) int {
	helper := make([]int, r-l+1)
	i := 0 // 对应helper的下标
	p1 := l
	p2 := m + 1
	ans := 0 // 左边比右边大的数字组合

	for p1 <= m && p2 <= r {
		if nums[p1] > nums[p2] {
			ans += m - p1 + 1 // 因为左边有序，所以数量 = 中点-p1+1
			helper[i] = nums[p2]
			p2++
		} else {
			helper[i] = nums[p1]
			p1++
		}
		i++
	}

	for p1 <= m {
		helper[i] = nums[p1]
		p1++
		i++
	}
	for p2 <= r {
		helper[i] = nums[p2]
		p2++
		i++
	}

	for i := 0; i < len(helper); i++ {
		nums[l+i] = helper[i]
	}

	return ans
}
