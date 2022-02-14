package reverselinkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "t1",
			args: args{
				head: genLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: genLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "t2",
			args: args{
				head: genLinkedList([]int{1, 2}),
			},
			want: genLinkedList([]int{2, 1}),
		},
		{
			name: "t3",
			args: args{
				head: genLinkedList([]int{}),
			},
			want: genLinkedList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func genLinkedList(vals []int) *ListNode {
	if len(vals) < 1 {
		return nil
	}

	var pre, head *ListNode = nil, nil
	for _, val := range vals {
		node := &ListNode{Val: val}
		if pre != nil {
			pre.Next = node
		} else {
			head = node
		}
		pre = node
	}

	return head
}
