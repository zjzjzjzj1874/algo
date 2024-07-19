package middle

// 面试题 16.10. 生存人数
// 给定 N 个人的出生年份和死亡年份，第 i 个人的出生年份为 birth[i]，死亡年份为 death[i]，实现一个方法以计算生存人数最多的年份。
//
// 你可以假设所有人都出生于 1900 年至 2000 年（含 1900 和 2000 ）之间。如果一个人在某一年的任意时期处于生存状态，那么他应该被纳入那一年的统计中。例如，生于 1908 年、死于 1909 年的人应当被列入 1908 年和 1909 年的计数。
//
// 如果有多个年份生存人数相同且均为最大值，输出其中最小的年份。
//
// 示例：
//
// 输入：
// birth = [1900, 1901, 1950]
// death = [1948, 1951, 2000]
// 输出： 1901
//
// 提示：
//
// 0 < birth.length == death.length <= 10000
// birth[i] <= death[i]

// 解题：查分数组 + 前缀和
func maxAliveYear(birth []int, death []int) int {
	helper := make([]int, 102) // 1900-2000，申明102个长度，处理数据边界，死亡会被统计到下一年
	const year = 1900
	for i := 0; i < len(birth); i++ {
		helper[birth[i]-year]++
		helper[death[i]-year+1]-- // 下一年开始减少生存人数：题意要求
	}

	maxAlive := 0
	curAlive := 0
	ans := 0
	for i := range helper {
		// 求前缀和 = 因为
		curAlive += helper[i]
		if curAlive > maxAlive {
			maxAlive = curAlive
			ans = i + year
		}
	}

	return ans
}
