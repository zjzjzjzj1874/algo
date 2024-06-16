package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			slow++
		}
		fast++
	}

	return slow
}

func removeDuplicatesV2(nums []int) (ans int) {
	for left, right := 0, 0; right < len(nums); left = right {
		nums[ans] = nums[left]
		ans += 1
		for right < len(nums) && nums[right] == nums[left] {
			right++
		}
	}
	return ans
}
