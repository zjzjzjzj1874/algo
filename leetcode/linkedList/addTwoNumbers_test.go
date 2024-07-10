package linkedList

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "add2Numbers", args: args{
			l1: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		}, want: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		}},
		{name: "add2Numbers", args: args{
			l1: &ListNode{
				Val:  5,
				Next: nil,
			},
			l2: &ListNode{
				Val:  5,
				Next: nil,
			},
		}, want: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  0,
				Next: nil,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersWithStack(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
