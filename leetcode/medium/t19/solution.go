package t19

import . "github.com/techoc/leetcodego/leetcode/utils/listNode"

// 19.删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy //双指针
	for i := 0; i < n; i++ {     //移动 first 指针 n 步
		first = first.Next
	}
	for ; first != nil; first = first.Next { // 移动 first 和 second 指针到链表尾部
		second = second.Next
	}
	second.Next = second.Next.Next // 由于已经使用了虚拟头节点，所以可以直接删除 second.Next
	return dummy.Next
}
