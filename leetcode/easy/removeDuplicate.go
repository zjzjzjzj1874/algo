package easy

// 26. 删除有序数组中的重复项
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
// 返回 k 。
// 判题标准:
// 系统会用下面的代码来测试你的题解:
// int[] nums = [...]; // 输入数组
// int[] expectedNums = [...]; // 长度正确的期望答案
// int k = removeDuplicates(nums); // 调用
// assert k == expectedNums.length;
//
//	for (int i = 0; i < k; i++) {
//	   assert nums[i] == expectedNums[i];
//	}
//
// 如果所有断言都通过，那么您的题解将被 通过。
//
// 示例 1：
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2,_]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
// 示例 2：
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

// 解题(快慢指针)：定义两个指针 fast 和 slow 分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，慢指针表示下一个不同元素要填入的下标位置，初始时两个指针都指向下标 1。
// 假设数组 nums 的长度为 n。将快指针 fast依次遍历从 1 到 n−1 的每个位置，对于每个位置，如果 nums[fast]≠nums[fast−1]，
// 说明 nums[fast]和之前的元素都不同，因此将 nums[fast]的值复制到 nums[slow]，然后将 slow的值加 1，即指向下一个位置。
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slow, fast := 1, 1
	for fast < len(nums) {
		if nums[slow-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		//if nums[fast] != nums[fast-1] {
		//	nums[slow] = nums[fast]
		//	slow++
		//}
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

// 27. 原地移除元素
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
// 假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
// 更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
// 返回 k。
// 示例 1：
// 输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2,_,_]
// 解释：你的函数函数应该返回 k = 2, 并且 nums 中的前两个元素均为 2。
// 你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。
// 示例 2：
// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
// 输出：5, nums = [0,1,4,0,3,_,_,_]
// 解释：你的函数应该返回 k = 5，并且 nums 中的前五个元素为 0,0,1,3,4。
// 注意这五个元素可以任意顺序返回。
// 你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。

// 解题：因为可以不计顺序返回，所以使用双指针来做这个题：
// 一个指针在左，一个指针在右，左指针一个一个移动，如果左指针的值=目标值，
// 则把右指针的值拷贝到左指针，右指针的值--，本次循环中左指针的值不变；直到两个指针相遇
func removeElement(nums []int, val int) int {
	res := len(nums)

	for i := 0; i < res; {
		if val != nums[i] {
			i++
			continue
		}

		nums[i], nums[res-1] = nums[res-1], nums[i] // 交换位置，换到最后面，然后res--
		res--
	}

	return res
}
