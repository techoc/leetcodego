package t876

import . "github.com/techoc/leetcodego/leetcode/utils"

// 876.链表的中间结点
// https://leetcode.cn/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	// 快慢指针 快指针走2步 慢指针走1步
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
