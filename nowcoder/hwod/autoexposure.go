package hwod

// 简单的自动曝光
// 题目描述
// 一个图像有 n 个像素点，存储在一个长度为 n 的数组 img 里，每个像素点的取值范围[0,255] 的正整数。
// 请你给图像每个像素点值加上一个整数 k(可以是负数)，得到新图 newImg，使得新图 newImg 的所有像素平均值最接近中位值 128。
// 请输出这个整数 k。 输入描述
// n 个整数，中间用空格分开 例如:
// 0000
// 4 个数值，中间用空格分开
// 输出描述 一个整数 k
func autoExposure(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	total := n * 128
	for _, num := range nums {
		total -= num
	}

	return total / n
}
