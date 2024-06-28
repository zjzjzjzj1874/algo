package easy

import "math"

// leetocde-121:买卖股票的最佳时机
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
// 示例 1：
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//	注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 示例 2：
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
// 提示：
// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104

// 解题：1、 双指针，最指针找出尽可能小的值，右指针找出尽可能大的值，记录其差值，然后左右移动指针；
// 移动规则，左边比右边大，则移动左边；否则移动右边
//
//	func maxProfit(prices []int) int {
//		if len(prices) < 2 { // 不足两天，无法完成买卖
//			return 0
//		}
//
//		min := math.MaxInt
//		max := 0
//
//		for i := 0; i < len(prices); i++ {
//			if prices[i] < min {
//				min = prices[i]
//			} else {
//				if prices[i]-min > max {
//					max = prices[i] - min
//				}
//			}
//		}
//
//		return max
//	}
func maxProfit(prices []int) int {
	var (
		min = math.MaxInt
		max = 0
	)

	for _, p := range prices {
		if min > p {
			min = p
		} else {
			if max < p-min {
				max = p - min
			}
		}
	}

	return max
}
