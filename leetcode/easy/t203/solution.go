package t203

import . "github.com/techoc/leetcodego/leetcode/utils"

// 203. 移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	//递归调用 如果链表的元素等于传入的值 则返回链表的下一个节点
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
