package t23

import (
	"reflect"
	"testing"
)
import . "github.com/techoc/leetcodego/leetcode/utils"

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			name: "Case 1",
			args: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name: "Case 2",
			args: [][]int{},
			want: []int{},
		},
		{
			name: "Case 3",
			args: [][]int{{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := make([]*ListNode, len(tt.args))
			for i, v := range tt.args {
				args[i] = ListToListNode(v)
			}
			got := ListNodeToList(mergeKLists(args))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
