package easy

import (
	"reflect"
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{name: "addBinary", args: args{a: "11", b: "1"}, want: "100"},
		//{name: "addBinary", args: args{a: "1011", b: "1010"}, want: "10101"},
		{name: "addBinary", args: args{a: "1", b: "111"}, want: "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relativeSortArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "relativeSortArray", args: args{arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, arr2: []int{2, 1, 4, 3, 9, 6}}, want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
