package base

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	arr1 := []int{7, 2, 3, 8, 4, 5, 6, 1}
	//fmt.Println(mergeSort(arr1, 0, len(arr1)-1))
	fmt.Println("Merge = ", Merge(arr1, 0, len(arr1)-1))

	arr2 := []int{7, 1, 8, 8, 4, 9, 6, 1}
	fmt.Println(mergeSort(arr2, 0, len(arr2)-1))
}

func Merge(arr []int, begin, end int) []int {
	if begin == end {
		return arr
	}

	middle := begin + (end-begin)>>1 // 求两数中间值

	Merge(arr, begin, middle)
	Merge(arr, middle+1, end)

	MergeCore(arr, begin, middle, end)

	return arr
}

func MergeCore(arr []int, begin, middle, end int) []int {
	temp := make([]int, end-begin+1)
	i := 0           // 临时数组的角标
	p1 := begin      // 前半部分的指针
	p2 := middle + 1 // 后半部分的指针

	for p1 <= middle && p2 <= end {
		if arr[p1] > arr[p2] { // 大于交换，按从小打大排序
			temp[i] = arr[p2]
			i++
			p2++
		} else {
			temp[i] = arr[p1]
			i++
			p1++
		}
	}

	for p1 <= middle { // p2先跑完，把p1剩下的按顺序塞入temp中
		temp[i] = arr[p1]
		i++
		p1++
	}

	for p2 <= end { // p1先跑完，把p2剩下的塞入temp中
		temp[i] = arr[p2]
		i++
		p2++
	}

	for i := 0; i < len(temp); i++ {
		arr[begin+i] = temp[i]
	}

	return arr
}
