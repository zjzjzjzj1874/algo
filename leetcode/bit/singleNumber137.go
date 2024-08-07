package bit

// 137. 只出现一次的数字 II
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
// 你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。
//
// 示例 1：
//
// 输入：nums = [2,2,3,2]
// 输出：3
// 示例 2：
//
// 输入：nums = [0,1,0,1,0,1,99]
// 输出：99
//
// 提示：
//
// 1 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

// 解题：hash吧
func singleNumber137(nums []int) int {
	nMap := make(map[int]int) // map[num]count

	for i := range nums {
		nMap[nums[i]]++
	}

	for val, count := range nMap {
		if count == 1 {
			return val
		}
	}

	return -1
}

func singleNumber137WithBit(nums []int) int {
	ans := int32(0) // 最后要求的数字
	for i := 0; i < 32; i++ {
		carry := int32(0)
		for _, num := range nums {
			carry += int32(num) >> i & 1 // 统计[0-32)上出现0，1的次数
		}

		ans += carry % 3 << i // 还原对应I的数字
	}

	return int(ans)
}
