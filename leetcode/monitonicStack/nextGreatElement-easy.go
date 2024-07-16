package monitonicStack

// 496. 下一个更大元素 I
// nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
// 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
// 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
//
// 示例 1：
//
// 输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
// 输出：[-1,3,-1]
// 解释：nums1 中每个值的下一个更大元素如下所述：
// - 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
// - 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
// - 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
// 示例 2：
//
// 输入：nums1 = [2,4], nums2 = [1,2,3,4].
// 输出：[3,-1]
// 解释：nums1 中每个值的下一个更大元素如下所述：
// - 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
// - 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。
//
// 提示：
//
// 1 <= nums1.length <= nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 104
// nums1和nums2中所有整数 互不相同
// nums1 中的所有整数同样出现在 nums2 中
//
// 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？

// 解题：Golang使用切片模拟单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n2 := getNums2NextGreatEle(nums2) // 计算出来nums2的下一个更大的元素，
	n1Map := make(map[int]int)        // map[n1Val]n2Index

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				n1Map[nums1[i]] = j
				break
			}
		}
	}

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		ans[i] = n2[n1Map[num]]
	}

	return ans
}

func getNums2NextGreatEle(nums2 []int) []int {
	n := len(nums2)
	ans := make([]int, n)
	for i := range nums2 {
		ans[i] = -1
	}

	stack := make([]int, 0, n) // 单调栈，存储的是元素的角标

	for i := 0; i < n; i++ {
		// 单调栈有数据，并且i代表的元素比单调栈最顶端的数大，此时需要弹出最顶端元素
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1] // 移除顶端元素，继续比较
		}

		stack = append(stack, i)
	}

	return ans
}

func nextGreaterElementWithMap(nums1 []int, nums2 []int) []int {
	n2Map := getNums2NextGreatEleWithMap(nums2) // 计算出来nums2的下一个更大的元素，

	ans := make([]int, len(nums1))
	for i, num := range nums1 {
		ans[i] = n2Map[num]
	}

	return ans
}

func getNums2NextGreatEleWithMap(nums2 []int) map[int]int {
	n := len(nums2)
	mmp := make(map[int]int) // map[nums2.Val]nextGreatVal
	for i := range nums2 {
		mmp[nums2[i]] = -1
	}

	stack := make([]int, 0, n) // 单调栈，存储的是元素的角标

	for i := 0; i < n; i++ {
		// 单调栈有数据，并且i代表的元素比单调栈最顶端的数大，此时需要弹出最顶端元素
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			mmp[nums2[stack[len(stack)-1]]] = nums2[i]
			stack = stack[:len(stack)-1] // 移除顶端元素，继续比较
		}

		stack = append(stack, i)
	}

	return mmp
}
