package easy

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Printf("nums = %v,res = %d\n", nums, removeDuplicates(nums))

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("nums = %v,res = %d\n", nums, removeDuplicates(nums))
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "foo", args: args{nums: []int{0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 3, 4, 4}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v, Nums:%v", got, tt.want, tt.args.nums)
			}
		})
	}
}
