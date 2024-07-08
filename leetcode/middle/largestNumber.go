package middle

import (
	"fmt"
	"sort"
)

// 179. 最大数
// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
// 示例 1：
//
// 输入：nums = [10,2]
// 输出："210"
// 示例 2：
//
// 输入：nums = [3,30,34,5,9]
// 输出："9534330"
//
// 提示：
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 109

// 解题：相邻两数按照拼接之后大小排序，最后再遍历数组，使用string组合
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := nums[i]
		y := nums[j]
		sx := 10 // 拼接之后的是两位数，还是三位数，还是四位数
		sy := 10

		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}

		return sy*x+y > sx*y+x
	})
	// 按照拼接之后从大到小排序了
	if nums[0] == 0 {
		return "0"
	}

	ans := ""
	for _, num := range nums {
		ans = fmt.Sprintf("%s%d", ans, num)
	}

	return ans
}
