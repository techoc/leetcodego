package t225

type MyStack struct {
	queue []int
	size  int
}

func Constructor() MyStack {
	// 初始化一个数组
	return MyStack{queue: make([]int, 5)}
}

func (this *MyStack) Push(x int) {
	if this.size == len(this.queue)-1 {
		//扩容一倍于队列的长度 并用0填充
		this.queue = append(this.queue, 0)
	}
	//入栈
	this.size++
	this.queue[this.size] = x
}

func (this *MyStack) Pop() int {
	//出栈
	num := this.queue[this.size]
	this.size--
	return num
}

func (this *MyStack) Top() int {
	//返回栈顶元素
	return this.queue[this.size]
}

func (this *MyStack) Empty() bool {
	return this.size <= 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
