package t19

import (
	. "github.com/techoc/leetcodego/leetcode/utils/listNode"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name   string
		head   *ListNode
		n      int
		expect *ListNode
	}{
		{
			name:   "Case 1",
			head:   ListToListNode([]int{1, 2, 3, 4, 5}),
			n:      2,
			expect: ListToListNode([]int{1, 2, 3, 5}),
		},
		{
			name:   "Case 2",
			head:   ListToListNode([]int{1}),
			n:      1,
			expect: nil,
		},
		{
			name:   "Case 3",
			head:   ListToListNode([]int{1, 2}),
			n:      1,
			expect: ListToListNode([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.expect)
			}
		})
	}
}
