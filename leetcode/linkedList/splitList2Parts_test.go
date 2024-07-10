package linkedList

import (
	"reflect"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []*ListNode
	}{
		{name: "split", args: args{head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}, k: 5}, wantAns: []*ListNode{{Val: 1, Next: nil}, {Val: 2, Next: nil}, {Val: 3, Next: nil}, {Next: nil}, {Next: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("splitListToParts() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
