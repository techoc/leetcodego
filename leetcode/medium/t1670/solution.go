package t1670

import "container/list"

// FrontMiddleBackQueue
// https://leetcode.cn/problems/design-front-middle-back-queue
type FrontMiddleBackQueue struct {
	// left 表示左边
	left *list.List
	// right 表示右边
	right *list.List
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left:  list.New(),
		right: list.New(),
	}
}

// PushFront 头部插入
func (this *FrontMiddleBackQueue) PushFront(val int) {
	// 调用左边的头部插入方法
	this.left.PushFront(val)
	// 如果左子串的长度大于右子串的长度+2
	// 重新平衡两个子串
	if this.left.Len() == this.right.Len()+2 {
		this.right.PushFront(this.left.Back().Value.(int))
		this.left.Remove(this.left.Back())
	}
}

// PushMiddle 中间插入
func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	// 如果左子串的长度大于右子串的长度+1
	if this.left.Len() == this.right.Len()+1 {
		this.right.PushFront(this.left.Back().Value.(int))
		this.left.Remove(this.left.Back())
	}
	this.left.PushBack(val)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.right.PushBack(val)
	if this.left.Len()+1 == this.right.Len() {
		this.left.PushBack(this.right.Front().Value.(int))
		this.right.Remove(this.right.Front())
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.left.Len() == 0 {
		return -1
	}
	val := this.left.Front().Value.(int)
	this.left.Remove(this.left.Front())
	if this.left.Len()+1 == this.right.Len() {
		this.left.PushBack(this.right.Front().Value.(int))
		this.right.Remove(this.right.Front())
	}
	return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.left.Len() == 0 {
		return -1
	}
	val := this.left.Back().Value.(int)
	this.left.Remove(this.left.Back())
	if this.left.Len()+1 == this.right.Len() {
		this.left.PushBack(this.right.Front().Value.(int))
		this.right.Remove(this.right.Front())
	}
	return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.left.Len() == 0 {
		return -1
	}
	if this.right.Len() == 0 {
		val := this.left.Back().Value.(int)
		this.left.Remove(this.left.Back())
		return val
	} else {
		val := this.right.Back().Value.(int)
		this.right.Remove(this.right.Back())
		if this.left.Len() == this.right.Len()+2 {
			this.right.PushFront(this.left.Back().Value.(int))
			this.left.Remove(this.left.Back())
		}
		return val
	}
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
