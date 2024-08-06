package math

import (
	"math/rand"
)

// 470. 用 Rand7() 实现 Rand10()
// 给定方法 rand7 可生成 [1,7] 范围内的均匀随机整数，试写一个方法 rand10 生成 [1,10] 范围内的均匀随机整数。
//
// 你只能调用 rand7() 且不能调用其他方法。请不要使用系统的 Math.random() 方法。
//
// 每个测试用例将有一个内部参数 n，即你实现的函数 rand10() 在测试时将被调用的次数。请注意，这不是传递给 rand10() 的参数。
//
// 示例 1:
//
// 输入: 1
// 输出: [2]
// 示例 2:
//
// 输入: 2
// 输出: [2,8]
// 示例 3:
//
// 输入: 3
// 输出: [3,8,10]
//
// 提示:
//
// 1 <= n <= 105
//
// 进阶:
//
// rand7()调用次数的 期望值 是多少 ?
// 你能否尽量少调用 rand7() ?

// 解题：使用01概率发生器，然后使用位运算来计算，当然，优化的方法也有很多，参考(https://leetcode.cn/problems/implement-rand10-using-rand7/solutions/427572/cong-pao-ying-bi-kai-shi-xun-xu-jian-jin-ba-zhe-da/)

func rand10() (ans int) {
	var gen01 func() int
	gen01 = func() int {
		for {
			gen := rand7()
			if gen == 4 {
				continue
			}
			if gen < 4 {
				return 0
			}
			if gen > 4 {
				return 1
			}
		}
	}

	//[1,10] => [0,9]+1

	for {
		ans = gen01()<<3 + gen01()<<2 + gen01()<<1 + gen01()
		if ans > 9 {
			continue
		} else {
			ans += 1
			break
		}
	}

	return
}

func rand7() int {
	return rand.Intn(7) + 1
}
