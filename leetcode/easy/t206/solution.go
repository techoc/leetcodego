package t206

import . "github.com/techoc/leetcodego/leetcode/utils"

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list
func reverseList(head *ListNode) *ListNode {
	//存储上一个节点
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
