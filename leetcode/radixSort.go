package leetcode

import "fmt"

// 基数排序
const radix = 10 // 基数

func RadixSort(nums []int) {
	fmt.Println("排序前：", nums)

	n := len(nums)
	if n < 2 {
		return
	}
	max := nums[0] // 找出最大值
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	helpBucket := make([]int, len(nums)) // 准备nums个数的辅助空间
	for d := 1; max/d > 0; d *= 10 {     // 表示有几位数，就进出多少次桶；100有三位数，进出桶三次
		count := make([]int, radix) // 0-9有10位，准备10个空间即可

		for i := 0; i < len(nums); i++ {
			j := (nums[i] / d) % 10 // 对应位数上的值
			count[j]++
		}

		// 更改计数[i]，使其包含该数字在输出中的实际位置
		for i := 1; i < radix; i++ {
			count[i] += count[i-1]
		}

		for i := len(nums) - 1; i >= 0; i-- {
			j := (nums[i] / d) % 10 // 对应位数上的值
			helpBucket[count[j]-1] = nums[i]
			count[j]--
		}

		// 将输出数组复制到arr，使arr现在包含根据当前数字排序的数字
		for i := 0; i < len(nums); i++ {
			nums[i] = helpBucket[i]
		}
	}

	fmt.Println("排序后：", nums)
}
