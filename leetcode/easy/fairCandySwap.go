package easy

// 888. 公平的糖果交换
// 爱丽丝和鲍勃拥有不同总数量的糖果。给你两个数组 aliceSizes 和 bobSizes ，aliceSizes[i] 是爱丽丝拥有的第 i 盒糖果中的糖果数量，bobSizes[j] 是鲍勃拥有的第 j 盒糖果中的糖果数量。
//
// 两人想要互相交换一盒糖果，这样在交换之后，他们就可以拥有相同总数量的糖果。一个人拥有的糖果总数量是他们每盒糖果数量的总和。
//
// 返回一个整数数组 answer，其中 answer[0] 是爱丽丝必须交换的糖果盒中的糖果的数目，answer[1] 是鲍勃必须交换的糖果盒中的糖果的数目。如果存在多个答案，你可以返回其中 任何一个 。题目测试用例保证存在与输入对应的答案。
//
// 示例 1：
//
// 输入：aliceSizes = [1,1], bobSizes = [2,2]
// 输出：[1,2]
// 示例 2：
//
// 输入：aliceSizes = [1,2], bobSizes = [2,3]
// 输出：[1,2]
// 示例 3：
//
// 输入：aliceSizes = [2], bobSizes = [1,3]
// 输出：[2,3]
// 示例 4：
//
// 输入：aliceSizes = [1,2,5], bobSizes = [2,4]
// 输出：[5,4]
//
// 提示：
//
// 1 <= aliceSizes.length, bobSizes.length <= 104
// 1 <= aliceSizes[i], bobSizes[j] <= 105
// 爱丽丝和鲍勃的糖果总数量不同。
// 题目数据保证对于给定的输入至少存在一个有效答案。

// 解题思路：想让糖果数量一致，则先计算差值，sa-sb = diff；多的人和少的那个人交换diff和diff/2的糖果即可。或者交换差值为diff/2的元素。
// 解题一：时间O(N^2)，空间：O(1)
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sa, sb := 0, 0

	for i := range aliceSizes {
		sa += aliceSizes[i]
	}
	for i := range bobSizes {
		sb += bobSizes[i]
	}

	diff := 0
	if sa > sb {
		diff = sa - sb
		for _, na := range aliceSizes {
			for _, nb := range bobSizes {
				if na-nb == diff/2 {
					return []int{na, nb}
				}
			}
		}
	} else {
		diff = sb - sa

		for _, na := range aliceSizes {
			for _, nb := range bobSizes {
				if nb-na == diff/2 {
					return []int{na, nb}
				}
			}
		}
	}

	return []int{}
}

// 解题：空间换时间，加速
func fairCandySwapOn(aliceSizes []int, bobSizes []int) []int {
	sa, sb := 0, 0
	aMap := make(map[int]struct{})
	for i := range aliceSizes {
		sa += aliceSizes[i]
		aMap[aliceSizes[i]] = struct{}{}
	}
	for i := range bobSizes {
		sb += bobSizes[i]
	}

	// answer[0] 是爱丽丝必须交换的糖果盒中的糖果的数目
	if sa > sb { // alice糖果多，出大的
		//sa - sb = diff; na-nb =diff/2
		diff := sa - sb
		for i := range bobSizes {
			_, ok := aMap[bobSizes[i]+diff/2]
			if ok {
				return []int{bobSizes[i] + diff/2, bobSizes[i]}
			}
		}
	}
	diff := sb - sa // bob多
	for i := range bobSizes {
		_, ok := aMap[bobSizes[i]-diff/2]
		if ok {
			return []int{bobSizes[i] - diff/2, bobSizes[i]}
		}
	}

	return []int{}
}

// 解题：空间换时间，加速
func fairCandySwapWithLT(aliceSizes []int, bobSizes []int) []int {
	sa, sb := 0, 0
	aMap := make(map[int]struct{})
	for i := range aliceSizes {
		sa += aliceSizes[i]
		aMap[aliceSizes[i]] = struct{}{}
	}
	for i := range bobSizes {
		sb += bobSizes[i]
	}

	delta := (sa - sb) / 2 // 糖果的差值
	// answer[0] 是爱丽丝必须交换的糖果盒中的糖果的数目
	for _, num := range bobSizes {
		if _, ok := aMap[num+delta]; ok {
			return []int{num + delta, num}
		}
	}

	return []int{}
}
