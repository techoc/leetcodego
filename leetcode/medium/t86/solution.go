package t86

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 86. 分隔链表
// https://leetcode.cn/problems/partition-list
func partition(head *ListNode, x int) *ListNode {
	// 创建两个链表，分别存储小于x和等于x的节点
	less := &ListNode{}
	more := &ListNode{}
	// 创建两个链表的虚拟头节点 方便插入节点
	lessHead := less
	moreHead := more
	// 遍历链表，将节点插入对应的链表
	for head != nil {
		if head.Val < x {
			// 小于 x 的值放入 less 链表
			less.Next = head
			less = less.Next
			head = head.Next
			// 断开 head 和 less 的连接
			less.Next = nil
		} else {
			// 大于等于 x 的值放入 more 链表
			more.Next = head
			more = more.Next
			head = head.Next
			// 断开 head 和 more 的连接
			more.Next = nil
		}
	}
	// 使用初始化的虚拟头节点 将 more 链表连接在 less 链表之后
	less.Next = moreHead.Next
	// 返回 less 链表虚拟头节点的下一个节点
	return lessHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
