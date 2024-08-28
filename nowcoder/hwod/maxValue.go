package hwod

import "sort"

// 双十一
// 题目描述
// 双十一众多商品进行打折销售，小明想购买一些自己心仪的商品，
// 但由于受购买资金限制，所以他决定从众多心意商品中购买3件，
// 而且想尽可能的花完资金，
// 现在请你设计一个程序帮助小明计算尽可能花费的最大资金额。
// 输入描述
// 第一行为整型数组M，数组长度小于100，数组元素记录单个商品的价格；
// 单个商品价格小于1000；
// 第二行输入为购买资金的额度R；
// R < 100000。
// ¶输出描述
// 输出为满足上述条件的最大花费额度
// 如果不存在满足上述条件的商品请返回-1
// 示例一
// 输入
// 23,26,36,27
// 78
// 输出
// 76
// 示例二
// 输入
// 23,30,40
// 26
// 输出
// -1
func maxValue(prices []int, money int) (ans int) {
	n := len(prices)
	sort.Ints(prices) // 对price排序
	if n < 3 || prices[0]+prices[1]+prices[2] > money {
		return -1
	}

	for i := 0; i < len(prices)-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			tmp := prices[i] + prices[l] + prices[r]
			if tmp > money {
				r--
			} else if tmp == money {
				return money
			} else {
				ans = max(ans, tmp)
				l++
			}
		}
	}

	return
}
