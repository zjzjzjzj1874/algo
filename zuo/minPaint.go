package zuo

import "fmt"

// 最小染色
// 牛牛有一些排成一行的正方形。每个正方形已经被染成红色或者绿色。牛牛现在可以选择任意一个正方形然后用这两种颜色的任意一种进行染色，这个正方形的颜色将
// 会被覆盖。牛牛的目标是在完成染色之后，每个红色R都比每个绿色G距离最左侧近。牛牛想知道他最少需要涂染几个正方形。如样例所示：s=RGRGR我们涂染之后变成RRRGG满足要求了，涂染的个数为2，没有比这个更好的涂染方案。

// 解题：动态规划，从左到右，依次染成一种颜色，然后依次记录需要涂染个数，时间O(N^2)，空间O(N);使用一个变量记录最小的也可以，从O(N)变成O(1)
func minPaint(s string) (ans int) {
	// 1、从0开始，左边全是R，右边全是绿色，则 ans = 右边全部染成绿色需要的个数
	// 2、从n-1开始，左边全是R，右边全是绿色，则 ans = 左边全部染成红色需要的个数
	// 3、0-n-1之间，ans = 左边染成红色，右边染成绿色的个数

	const (
		R = 'R' // 红色
		G = 'G' // 绿色
	)
	if s == "" {
		// 特殊边界条件
		return
	}
	n := len(s)
	for i := 0; i <= n; i++ {
		tmp := 0
		ts := ""

		// 左边染成红色
		for j := 0; j < i; j++ {
			ts += string(R)
			if s[j] != R {
				tmp++
			}
		}

		// 右边染成绿色
		for j := i; j < n; j++ {
			ts += string(G)

			if s[j] != G {
				tmp++
			}
		}
		fmt.Println("改变后：", ts, ",长度：", tmp)
		if ans == 0 || (tmp != 0 && ans != 0 && tmp < ans) {
			ans = tmp
		}
	}

	return
}

// 解题：数组预处理，从 O(N^2) 到 O(N)；
func minPaintWithON(s string) (ans int) {
	// 1、从0开始，左边全是R，右边全是绿色，则 ans = 右边全部染成绿色需要的个数
	// 2、从n-1开始，左边全是R，右边全是绿色，则 ans = 左边全部染成红色需要的个数
	// 3、0-n-1之间，ans = 左边染成红色，右边染成绿色的个数

	const (
		R = 'R' // 红色
		G = 'G' // 绿色
	)
	if s == "" {
		// 特殊边界条件
		return
	}
	n := len(s)
	arrG := make([]int, n) // 从左往右，[0,i]的位置上有多少个G
	arrR := make([]int, n) // 从右往左，[0,i]的位置上有多少个R
	for i := 0; i < n; i++ {
		if s[i] == G {
			arrG[i]++
		}
		if s[n-1-i] == R {
			arrR[n-1-i]++
		}
		if i != 0 {
			arrG[i] += arrG[i-1]
			arrR[n-1-i] += arrR[n-i]
		}

	}

	for i := 0; i <= n; i++ {
		tmp := 0
		if i == 0 {
			tmp = arrR[i]
		} else if i == n {
			tmp = arrG[i-1]
		} else {
			tmp = arrR[i] + arrG[i-1]
		}
		if ans == 0 || tmp < ans {
			ans = tmp
		}
	}

	return
}
