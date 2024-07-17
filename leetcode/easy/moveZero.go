package easy

// 283. 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 示例 2:
// 输入: nums = [0]
// 输出: [0]
// 提示:
// 1 <= nums.length <= 104
// -231 <= nums[i] <= 231 - 1
// 进阶：你能尽量减少完成的操作次数吗？

// 解题：双指针，快慢指针，快指针取遍历数组，慢指针保留最近一个非零元素的index,快慢指针中间的元素一定是0
//
//	func moveZeroes(nums []int) {
//		fast, slow := 0, 0
//
//		for fast < len(nums) {
//			if nums[fast] != 0 { // 快指针遇到0了，要和慢指针的交换位置
//				nums[fast], nums[slow] = nums[slow], nums[fast]
//				slow++
//			}
//			fast++
//		}
//	}
func moveZeroes(nums []int) {
	// 解题：维护两个指针，一个指针指向0，一个指向非0，然后非0和0遇到后交换位置

	left := -1 // 指向0的指针
	// right := 0 // 正常往右移动的指针

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			if left == -1 {
				left = right // 遇到0指针
			}
			continue
		}

		if left != -1 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}
}
