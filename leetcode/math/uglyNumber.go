package math

// 1201. 丑数 III
// 丑数是可以被 a 或 b 或 c 整除的 正整数 。
//
// 给你四个整数：n 、a 、b 、c ，请你设计一个算法来找出第 n 个丑数。
//
// 示例 1：
//
// 输入：n = 3, a = 2, b = 3, c = 5
// 输出：4
// 解释：丑数序列为 2, 3, 4, 5, 6, 8, 9, 10... 其中第 3 个是 4。
// 示例 2：
//
// 输入：n = 4, a = 2, b = 3, c = 4
// 输出：6
// 解释：丑数序列为 2, 3, 4, 6, 8, 9, 10, 12... 其中第 4 个是 6。
// 示例 3：
//
// 输入：n = 5, a = 2, b = 11, c = 13
// 输出：10
// 解释：丑数序列为 2, 4, 6, 8, 10, 11, 12, 13... 其中第 5 个是 10。
//
// 提示：
//
// 1 <= n, a, b, c <= 109
// 1 <= a * b * c <= 1018
// 本题结果在 [1, 2 * 109] 的范围内

// 二分法思路
// 思路如下：
// 1、我们先找到a,b,c里最小的那个数，比如是a，那么第n个丑数肯定是小于等于n * a的对不对？
// 因为 0 到 n*a范围内是有可能出现数字，可以被b或c整除的
// 2、然后就开始二分法的做法了，将n * a置为上限ceil，0置为下限0。
// 3、我们求解mid = (ceil+floor)/2这个数里包含了多少丑数，具体解法下面另外说。
// 4、如果上一步的数字等于n,那最好啦，判断当前的mid是否是丑数，如果是，直接返回mid，如果不是，将ceil置为mid - 1；
// 如果上一步的数字大于n,将ceil置为mid - 1；
// 如果上一步的数字小于n,将floor置为mid + 1；
// 为啥数字等于n，不能直接返回当前mid值呢？
// 比如：在a = 2, b = 21, c = 31的情况下，14和15这两个范围里都包含了7个丑数，如下所示
// 2 4 6 8 10 12 14
// 但是很明显，14才是第7个丑数。
// 求指定数字范围内的丑数数量
// 指定数字num范围内的丑数数量为：
// num/a + num/b + num/c - num/lcm(ab) - num/lcm(ac) -  num/lcm(bc) + num/lcm(abc)
func nthUglyNumber(n int, a int, b int, c int) int {
	if a == 1 || b == 1 || c == 1 {
		return n
	}

	ab := lcm(a, b)
	bc := lcm(b, c)
	ac := lcm(a, c)
	abc := lcm(ab, c)
	ceil := n * min(a, b, c)
	floor := 0
	for floor <= ceil {
		mid := floor + (ceil-floor)/2
		num := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if num == n {
			if mid%a == 0 || mid%b == 0 || mid%c == 0 {
				return mid
			} else {
				ceil = mid - 1
			}
		} else if num > n {
			ceil = mid - 1
		} else {
			floor = mid + 1
		}
	}

	return floor
}

// 264. 丑数 II
// 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
//
// 丑数 就是质因子只包含 2、3 和 5 的正整数。
//
// 示例 1：
//
// 输入：n = 10
// 输出：12
// 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
// 示例 2：
//
// 输入：n = 1
// 输出：1
// 解释：1 通常被视为丑数。
//
// 提示：
//
// 1 <= n <= 1690
func nthUglyNumberII(n int) int {
	if n == 1 {
		return 1
	}
	// 先用[]int{}的数组来装n；一会儿使用O(1)优化
	dp := make([]int, n+1)
	dp[1] = 1          // 初始化
	i, j, k := 1, 1, 1 // 分别表示对应2,3,5的指针

	for m := 2; m <= n; m++ {
		ia := dp[i] * 2
		jb := dp[j] * 3
		kc := dp[k] * 5
		dp[m] = min(ia, jb, kc)
		if dp[m] == ia {
			i++
		}
		if dp[m] == jb {
			j++
		}
		if dp[m] == kc {
			k++
		}
	}

	return dp[n]
}

// 解题：暴力尝试：O(N):会超时
func nthUglyNumberTimeout(n int, a int, b int, c int) int {
	if a == 1 || b == 1 || c == 1 {
		return n
	}

	// 先用[]int{}的数组来装n；一会儿使用O(1)优化
	i, j, k := 1, 1, 1 // 分别表示a, b, c的倍数，

	ans := 0
	for n > 0 {
		ia := i * a
		jb := j * b
		kc := k * c
		m := min(ia, jb, kc)
		if m == ia {
			i++
			ans = ia
		}
		if m == jb {
			j++
			ans = jb
		}
		if m == kc {
			k++
			ans = kc
		}
		n--
	}

	return ans
}
