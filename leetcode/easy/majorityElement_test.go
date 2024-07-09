package easy

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{name: "majority", want: []int{3}, args: args{nums: []int{3, 2, 3}}},
		{name: "majority", want: []int{1, 2}, args: args{nums: []int{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementWithMor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
