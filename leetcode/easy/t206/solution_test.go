package t206

import (
	. "github.com/techoc/leetcodego/leetcode/utils"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Case 1",
			head: ListToListNode([]int{1, 2, 3, 4, 5}),
			want: ListToListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "Case 2",
			head: ListToListNode([]int{1, 2}),
			want: ListToListNode([]int{2, 1}),
		},
		{
			name: "Case 3",
			head: ListToListNode([]int{}),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
