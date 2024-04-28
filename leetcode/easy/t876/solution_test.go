package t876

import (
	. "github.com/techoc/leetcodego/leetcode/utils"
	"reflect"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Case 1",
			head: ListToListNode([]int{1, 2, 3, 4, 5}),
			want: ListToListNode([]int{3, 4, 5}),
		},
		{
			name: "Case 2",
			head: ListToListNode([]int{1, 2, 3, 4, 5, 6}),
			want: ListToListNode([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
