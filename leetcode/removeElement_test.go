package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Printf("nums = %v,res = %d\n", nums, removeElement(nums, val))
}
