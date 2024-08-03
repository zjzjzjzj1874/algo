package dp

import "math"

// 1884. 鸡蛋掉落-两枚鸡蛋
// 给你 2 枚相同 的鸡蛋，和一栋从第 1 层到第 n 层共有 n 层楼的建筑。
//
// 已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都 会碎 ，从 f 楼层或比它低 的楼层落下的鸡蛋都 不会碎 。
//
// 每次操作，你可以取一枚 没有碎 的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。
//
// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
//
// 示例 1：
//
// 输入：n = 2
// 输出：2
// 解释：我们可以将第一枚鸡蛋从 1 楼扔下，然后将第二枚从 2 楼扔下。
// 如果第一枚鸡蛋碎了，可知 f = 0；
// 如果第二枚鸡蛋碎了，但第一枚没碎，可知 f = 1；
// 否则，当两个鸡蛋都没碎时，可知 f = 2。
// 示例 2：
//
// 输入：n = 100
// 输出：14
// 解释：
// 一种最优的策略是：
// - 将第一枚鸡蛋从 9 楼扔下。如果碎了，那么 f 在 0 和 8 之间。将第二枚从 1 楼扔下，然后每扔一次上一层楼，在 8 次内找到 f 。总操作次数 = 1 + 8 = 9 。
// - 如果第一枚鸡蛋没有碎，那么再把第一枚鸡蛋从 22 层扔下。如果碎了，那么 f 在 9 和 21 之间。将第二枚鸡蛋从 10 楼扔下，然后每扔一次上一层楼，在 12 次内找到 f 。总操作次数 = 2 + 12 = 14 。
// - 如果第一枚鸡蛋没有再次碎掉，则按照类似的方法从 34, 45, 55, 64, 72, 79, 85, 90, 94, 97, 99 和 100 楼分别扔下第一枚鸡蛋。
// 不管结果如何，最多需要扔 14 次来确定 f 。
//
// 提示：
//
// 1 <= n <= 1000
// 思路：先把n分成多少堆， 这个堆就是要找的次数；从后往前推，n到n-1需要验证一次，那么n-1到n-3，n-3,n-6... 越往前，step的长度就+1，抵消后面本身的长度

func twoEggDrop(n int) int {
	step := 1
	times := 0
	for i := n; i > 0; {
		i = i - step
		step++
		times++
	}

	return times
}

// 887. 鸡蛋掉落
// 太难了，战术性放弃 TODO 答案是不对的
func superEggDrop(k int, n int) (ans int) {
	// 分情况讨论，如果k=1,ans = n
	// k = 2,参考上一个鸡蛋掉落的问题，
	// 如果k>2，此时可以使用二分缩小范围
	// 如果2^k > n,直接返回log(n)

	if math.Pow(2, float64(k)) == float64(n) {
		ans += k
		return
	} else if math.Pow(2, float64(k)) > float64(n) {
		for math.Pow(2, float64(k)) >= float64(n) {
			k--
		}
		ans += k
		return
	}
	var egg func(k, n int)
	egg = func(k, n int) {
		if k == 1 {
			ans += n
			return
		}
		if k == 2 {
			ans += twoEggDrop(n)
			return
		}

		ans++
		egg(k-1, n/2)
	}

	egg(k, n)
	return
}
