package hwod

// 新员工座位安排系统
// 题目描述
// 工位由序列 F1,F2...Fn 组成，Fi 值为 0、1 或 2。其中 0 代表空置，1 代表有人，2 代表障碍物。 1、某一空位的友好度为左右连续老员工数之和 2、为方便新员工学习求助，
// 优先安排友好度高的空位 给出工位序列，求所有空位中友好度的最大值。
// 输入描述
// 第一行为工位序列:F1,F2...Fn 组成，1<=n<=100000，Fi 值为 0、1 或 2。其中 0 代表空置，1 代 码有人，2 代表障碍物
// 其中 0 代表空置，1 代码有人，2 代表障碍物。
// 输出描述 所有空位中友好度的最大值。如果没有空位，返回 0。
// 解题：两次遍历数组即可得到
func maxFriendship(seats []int) (ans int) {
	n := len(seats)
	left := make([]int, n)
	leftPeople, rightPeople := 0, 0

	// 从左往右数
	for i := 0; i < n; i++ {
		if seats[i] == 0 {
			ans = max(ans, leftPeople)
			left[i] = ans
			leftPeople = 0
		} else if seats[i] == 1 {
			left[i] = 0
			leftPeople++
		} else {
			left[i] = 0
			leftPeople = 0
		}
	}
	// 从右往左
	for i := n - 1; i >= 0; i-- {
		if seats[i] == 0 {
			ans = max(ans, rightPeople+left[i])
			rightPeople = 0
		} else if seats[i] == 1 {
			rightPeople++
		} else {
			rightPeople = 0
		}
	}

	return
}
