package leetcode

import (
	"fmt"
)

// HeapSort 堆排
func HeapSort(nums []int) {
	fmt.Println("排序前：", nums)

	n := len(nums)

	// 构建一个大根堆，为什么构建大根堆，因为把最大的数放在数组最后面
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	// 从堆中依次提取元素
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] // 把当前元素移动到末尾
		heapify(nums, i, 0)                 // 减少元素后继续调用堆化，把第一个元素放在合适的位置上
	}

	fmt.Println("排序后：", nums)
}

func heapInsert(nums []int, index int) {
	for nums[index] > nums[(index-1)/2] { // 如果当前节点大于父节点。
		nums[index], nums[(index-1)/2] = nums[(index-1)/2], nums[index]
		index = (index - 1) / 2 // 当前节点向上移动到父节点
	}
}

// 堆化：堆结构最重要的一个操作:将数组转换成堆结构，n表示长度，i表示下标
func heapify(nums []int, n, i int) {
	for {
		// 位置为i的数和左右孩子pk，谁最大谁胜出
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && nums[left] > nums[largest] {
			largest = left
		}

		if right < n && nums[right] > nums[largest] {
			largest = right
		}

		if largest == i { // 如果i比左右孩子都大，则停止循环
			break
		}

		// 否则，交换i和左右孩子中较大的数
		nums[i], nums[largest] = nums[largest], nums[i]
		i = largest
	}
}
