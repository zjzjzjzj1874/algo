package main

import "fmt"

//HJ41 称砝码
//描述
//现有n种砝码，重量互不相等，分别为 m1,m2,m3…mn ；
//每种砝码对应的数量为 x1,x2,x3...xn 。现在要用这些砝码去称物体的重量(放在同一侧)，问能称出多少种不同的重量。
//注：
//称重重量包括 0
//数据范围：每组输入数据满足1≤n≤10，1≤𝑚𝑖≤2000，1≤𝑥𝑖≤10
//输入描述：
//对于每组测试数据：
//第一行：n --- 砝码的种数(范围[1,10])
//第二行：m1 m2 m3 ... mn --- 每种砝码的重量(范围[1,2000])
//第三行：x1 x2 x3 .... xn --- 每种砝码对应的数量(范围[1,10])
//输出描述：
//利用给定的砝码可以称出的不同的重量数
//
//示例1
//输入：
//2
//1 2
//2 1
//复制
//输出：
//5
//复制
//说明：
//可以表示出0，1，2，3，4五种重量。

func main() {
	n := 0
	fmt.Scan(&n)
	m := make([]int, n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		mi := 0
		fmt.Scan(&mi)
		m[i] = mi
		nums[i] = mi
	}
	for i := 0; i < n; i++ {
		mi := 0
		fmt.Scan(&mi)
		if mi == 1 {
			continue
		}
		// 已经写入一个了
		for j := 1; j < mi; j++ {
			nums = append(nums, m[i])
		}
	}

	set := map[int]struct{}{0: {}}
	for i := range nums {
		tmp := map[int]struct{}{}
		for j := range set {
			tmp[nums[i]+j] = struct{}{}
		}
		for j := range tmp {
			set[j] = struct{}{}
		}
	}
	fmt.Println(len(set))
}
