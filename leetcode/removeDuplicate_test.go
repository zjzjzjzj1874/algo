package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Printf("nums = %v,res = %d\n", nums, removeDuplicates(nums))

	//nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums = []int{1, 1, 2, 4, 1, 1}
	fmt.Printf("nums = %v,res = %d\n", nums, removeDuplicatesV2(nums))
}
