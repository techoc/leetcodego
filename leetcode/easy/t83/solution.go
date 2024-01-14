package t83

// 83.删除排序链表中的重复元素
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
