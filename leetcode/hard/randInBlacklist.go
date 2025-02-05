/*
 * @Author: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @Date: 2025-01-14 13:56:53
 * @LastEditors: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @LastEditTime: 2025-01-14 14:10:15
 * @FilePath: /algo/leetcode/hard/randInBlacklist.go
 * @Description: 710 黑名单中的随机数
 */
package hard

import "math/rand"

// 710. 黑名单中的随机数
// 给定一个整数 n 和一个 无重复 黑名单整数数组 blacklist 。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个 未加入 黑名单 blacklist 的整数。任何在上述范围内且不在黑名单 blacklist 中的整数都应该有 同等的可能性 被返回。
// 优化你的算法，使它最小化调用语言 内置 随机函数的次数。
// 实现 Solution 类:
// Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
// int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数
// 示例 1：
// 输入
// ["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
// [[7, [2, 3, 5]], [], [], [], [], [], [], []]
// 输出
// [null, 0, 4, 1, 6, 1, 0, 4]
// 解释
// Solution solution = new Solution(7, [2, 3, 5]);
// solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
// 0、1、4和6的返回概率必须相等(即概率为1/4)。
// solution.pick(); // 返回 4
// solution.pick(); // 返回 1
// solution.pick(); // 返回 6
// solution.pick(); // 返回 1
// solution.pick(); // 返回 0
// solution.pick(); // 返回 4
// 提示:
// 1 <= n <= 109
// 0 <= blacklist.length <= min(105, n - 1)
// 0 <= blacklist[i] < n
// blacklist 中所有值都 不同
//  pick 最多被调用 2 * 104 次

type Solution struct {
	blackMap map[int]struct{}
	n        int
}

func Constructor(n int, blacklist []int) Solution {
	bm := make(map[int]struct{})
	for _, val := range blacklist {
		bm[val] = struct{}{}
	}

	return Solution{
		blackMap: bm,
		n:        n,
	}
}

func (this *Solution) Pick() int {
	num := rand.Intn(this.n)
	_, ok := this.blackMap[num]
	for ok {
		num = rand.Intn(this.n)
		_, ok = this.blackMap[num]
	}

	return num
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
