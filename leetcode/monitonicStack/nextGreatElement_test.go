package monitonicStack

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "nextGreatElements", args: args{nums: []int{1, 2, 1}}, want: []int{2, -1, 2}},
		{name: "nextGreatElements", args: args{nums: []int{1, 2, 3, 4, 3}}, want: []int{2, 3, 4, -1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "nextGreater", args: args{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}}, want: []int{-1, 3, -1}},
		{name: "nextGreater", args: args{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}}, want: []int{3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementWithMap(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
