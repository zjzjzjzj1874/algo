package base

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	arr1 := []int{7, 2, 3, 8, 4, 5, 6, 1}
	fmt.Println(mergeSort(arr1, 0, len(arr1)-1))

	arr2 := []int{7, 1, 8, 8, 4, 9, 6, 1}
	fmt.Println(mergeSort(arr2, 0, len(arr2)-1))
}
