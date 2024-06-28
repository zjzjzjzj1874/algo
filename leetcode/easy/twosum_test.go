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

func Test_twosumHash(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "twosum", args: args{arr: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twosumHash(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twosumHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
