package easy

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twosumTwoPtr(nums, target))

	//nums = []int{3, 2, 4}
	//target = 6
	//fmt.Println("twosumHash == ", twosumHash(nums, target))

	nums = []int{3, 3}
	target = 6
	bubbleRes := twosumBubble(nums, target)
	hashRes := twosumHash(nums, target)
	sort.Ints(bubbleRes)
	sort.Ints(hashRes)
	fmt.Printf("bubbleRes = %v, hashRes = %v,equal = %v", bubbleRes, hashRes, reflect.DeepEqual(bubbleRes, hashRes))
}
