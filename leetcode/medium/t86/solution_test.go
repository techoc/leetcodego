package t86

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		x    int
		want *ListNode
	}{
		{
			name: "Case 1",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 2,
								},
							},
						},
					},
				},
			},
			x: 3,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Case 2",
			head: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
			x: 2,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.head, tt.x); !reflect.DeepEqual(got, tt.want) {
				//if got := partition(tt.head, tt.x); !equals(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
