package prefixsum

// 303. 区域和检索 - 数组不可变
// 给定一个整数数组  nums，处理以下类型的多个查询:
//
// 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
// 实现 NumArray 类：
//
// NumArray(int[] nums) 使用数组 nums 初始化对象
// int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和 right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
//
// 示例 1：
//
// 输入：
// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// 输出：
// [null, 1, -1, -3]
//
// 解释：
// NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
// numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
// numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
// numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
//
// 提示：
//
// 1 <= nums.length <= 104
// -105 <= nums[i] <= 105
// 0 <= i <= j < nums.length
// 最多调用 10^4 次 sumRange 方法

type NumArray struct {
	nums      []int       // 原始数组
	prefixSum map[int]int // 前缀和
}

func Constructor(nums []int) NumArray {
	sum := 0
	ps := make(map[int]int)
	for i, num := range nums {
		sum += num
		ps[i] = sum
	}
	return NumArray{nums: nums, prefixSum: ps}
}

func (a *NumArray) SumRange(left int, right int) (ans int) {
	if left == 0 {
		return a.prefixSum[right]
	} else {
		return a.prefixSum[right] - a.prefixSum[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

//307. 区域和检索 - 数组可修改
//给你一个数组 nums ，请你完成两类查询。
//
//其中一类查询要求 更新 数组 nums 下标对应的值
//另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
//实现 NumArray 类：
//
//NumArray(int[] nums) 用整数数组 nums 初始化对象
//void update(int index, int val) 将 nums[index] 的值 更新 为 val
//int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
//
//
//示例 1：
//
//输入：
//["NumArray", "sumRange", "update", "sumRange"]
//[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
//输出：
//[null, 9, null, 8]
//
//解释：
//NumArray numArray = new NumArray([1, 3, 5]);
//numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
//numArray.update(1, 2);   // nums = [1,2,5]
//numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8
//
//
//提示：
//
//1 <= nums.length <= 3 * 104
//-100 <= nums[i] <= 100
//0 <= index < nums.length
//-100 <= val <= 100
//0 <= left <= right < nums.length
//调用 update 和 sumRange 方法次数不大于 3 * 104

// 鬼东西哦，要超时

func (a *NumArray) Update(index int, val int) {
	for i := index; i < len(a.nums); i++ {
		a.prefixSum[i] += val
	}
	a.nums[index] = val
}
