package middle

// 1093. 大样本统计
// 我们对 0 到 255 之间的整数进行采样，并将结果存储在数组 count 中：count[k] 就是整数 k 在样本中出现的次数。
//
// 计算以下统计数据:
//
// minimum ：样本中的最小元素。
// maximum ：样品中的最大元素。
// mean ：样本的平均值，计算为所有元素的总和除以元素总数。
// median ：
// 如果样本的元素个数是奇数，那么一旦样本排序后，中位数 median 就是中间的元素。
// 如果样本中有偶数个元素，那么中位数median 就是样本排序后中间两个元素的平均值。
// mode ：样本中出现次数最多的数字。保众数是 唯一 的。
// 以浮点数数组的形式返回样本的统计信息 [minimum, maximum, mean, median, mode] 。与真实答案误差在 10-5 内的答案都可以通过。
//
// 示例 1：
//
// 输入：count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// 输出：[1.00000,3.00000,2.37500,2.50000,3.00000]
// 解释：用count表示的样本为[1,2,2,2,3,3,3,3]。
// 最小值和最大值分别为1和3。
// 均值是(1+2+2+2+3+3+3+3) / 8 = 19 / 8 = 2.375。
// 因为样本的大小是偶数，所以中位数是中间两个元素2和3的平均值，也就是2.5。
// 众数为3，因为它在样本中出现的次数最多。
// 示例 2：
//
// 输入：count = [0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// 输出：[1.00000,4.00000,2.18182,2.00000,1.00000]
// 解释：用count表示的样本为[1,1,1,1,2,2,3,3,3,4,4]。
// 最小值为1，最大值为4。
// 平均数是(1+1+1+1+2+2+2+3+3+4+4)/ 11 = 24 / 11 = 2.18181818…(为了显示，输出显示了整数2.18182)。
// 因为样本的大小是奇数，所以中值是中间元素2。
// 众数为1，因为它在样本中出现的次数最多。
//
// 提示：
//
// count.length == 256
// 0 <= count[i] <= 109
// 1 <= sum(count) <= 109
// count 的众数是 唯一 的

// 解题思路：一次遍历，准备所有东西 ==> 超出内存限制
func sampleStatsOverMem(count []int) []float64 {
	ans := make([]float64, 5)
	ans[0] = -1

	mode := 0 // 出现最多次数
	modeNum := 0
	sum := 0      // 和
	numCount := 0 // 元素个数
	nums := make([]int, 0, 16)
	for i, cou := range count {
		if cou != 0 {
			numCount += cou // 元素个数
			sum += i * cou
			if ans[0] == -1 {
				ans[0] = float64(i) // 最小元素
			}

			ans[1] = float64(i) // 最大元素

			if cou > mode {
				mode = cou
				modeNum = i
			}

			for j := 0; j < cou; j++ {
				nums = append(nums, i)
			}
		}
	}

	ans[4] = float64(modeNum)                 // 众数
	ans[2] = float64(sum) / float64(numCount) // 平均数

	if numCount%2 == 0 { // 偶数
		ans[3] = float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / float64(2) // 奇数
	} else {
		ans[3] = float64(nums[len(nums)/2]) // 奇数
	}

	return ans
}

// 解题：优化，使用双指针

func sampleStats(count []int) []float64 {
	ans := make([]float64, 5)
	ans[0] = -1

	sum := 0      // 和
	numCount := 0 // 元素个数

	left := 0
	right := len(count) - 1
	maximum := -1  // 最大元素
	minimum := 256 // 最小元素
	mode := 0      // 出现最多次数
	modeNum := 0

	for left < right {
		l := count[left]
		r := count[right]
		if l != 0 {
			numCount += l // 元素个数
			sum += l * left
			if left > maximum {
				maximum = left
			}
			if left < minimum {
				minimum = left
			}
		}
		if r != 0 {
			numCount += r // 元素个数
			sum += r * right
			if right > maximum {
				maximum = right
			}
			if right < minimum {
				minimum = right
			}
		}

		if l > mode {
			mode = l
			modeNum = left
		}
		if r > mode {
			mode = r
			modeNum = right
		}

		left++
		right--
	}

	mid := numCount / 2
	//median
	median := float64(0) // 中位数
	if numCount%2 == 0 { // 偶数个
		next := -1
		for i, cou := range count {
			mid -= cou
			if mid == 0 {
				if next == -1 {
					median = float64(i)
					next = i
				}
			} else if mid < 0 {
				if next != -1 {
					// 说明从上一步来的，
					median = (median + float64(i)) / 2
				} else {
					median = float64(i)
				}
				break
			}
		}
	} else { // 奇数
		for i, cou := range count {
			mid -= cou
			if mid < 0 {
				median = float64(i)
				break
			}
		}
	}
	// 以浮点数数组的形式返回样本的统计信息 [minimum, maximum, mean, median, mode] 。与真实答案误差在 10-5 内的答案都可以通过。
	ans[0] = float64(minimum)                 // 最小
	ans[1] = float64(maximum)                 // 最大
	ans[2] = float64(sum) / float64(numCount) // 平均数
	ans[3] = median                           // 中位数
	ans[4] = float64(modeNum)                 // 众数

	return ans
}
