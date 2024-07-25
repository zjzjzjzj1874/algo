package linkedList

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		aNode *Node
		x     int
	}

	a := &Node{Val: 3}
	a.Next = &Node{Val: 3}
	a.Next.Next = &Node{Val: 3}
	a.Next.Next.Next = a
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "insert", args: args{aNode: a, x: 0}, want: &Node{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.aNode, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
