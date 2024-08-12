package easy

// 605. 种花问题
// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
// 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
//
// 示例 1：
//
// 输入：flowerbed = [1,0,0,0,1], n = 1
// 输出：true
// 示例 2：
//
// 输入：flowerbed = [1,0,0,0,1], n = 2
// 输出：false
//
// 提示：
//
// 1 <= flowerbed.length <= 2 * 104
// flowerbed[i] 为 0 或 1
// flowerbed 中不存在相邻的两朵花
// 0 <= n <= flowerbed.length
func canPlaceFlowers(flowerbed []int, n int) bool {
	m := len(flowerbed)
	if n > m {
		return false
	}
	if m == 1 && flowerbed[0] == 0 && n <= 1 {
		return true
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			if i+1 < m && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if i == m-1 {
			if i-1 >= 0 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else { // i [1, m-1]
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
	}

	return n <= 0
}

// 解题：额外数组，特殊处理0，n的位置
func canPlaceFlowersWithArr(flowerbed []int, n int) bool {
	m := len(flowerbed)
	if n > m {
		return false
	}

	// 把-1和n的位置用0补上
	flowerbed = append([]int{0}, append(flowerbed, 0)...)

	for i := 1; i <= m; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			n--
			i++ // 这个位置种花，下一个位置一定不能种花，跳过判断
		}
	}

	return n <= 0
}

// 解题：不用额外处理数组，原地数组
func canPlaceFlowersWithYuanDi(flowerbed []int, n int) bool {
	m := len(flowerbed)
	if n > m {
		return false
	}

	// 把-1和n的位置用0补上

	for i := 0; i < m; i++ {
		if (i == 0 || flowerbed[i-1] == 0) && flowerbed[i] == 0 && (i == m-1 || flowerbed[i+1] == 0) {
			n--
			i++ // 这个位置种花，下一个位置一定不能种花，跳过判断
		}
	}

	return n <= 0
}
