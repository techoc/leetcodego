package t215

// https://leetcode.cn/problems/kth-largest-element-in-an-array
// 215.数组中的第K个最大元素
// 由于题目要求其解法的时间复杂度为 O(n) ，所以可以选择使用堆排序或者快排
func findKthLargest(nums []int, k int) int {
	// 边界检查
	if k < 1 || k > len(nums) {
		return 0
	}
	// 堆排序 时间复杂度为O(nlogn)
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

// 创建大根堆
func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

// 最大堆化
// 1. 从堆顶开始，依次向下比较，将大的元素上浮到堆顶
// 2. 重复1步，直到堆的根节点
// 3. 最大堆化的时间复杂度为O(nlogn)
func maxHeapify(a []int, i, heapSize int) {
	// 定义左子树下标为i*2+1，右子树下标为i*2+2，最大子树下标为i
	l, r, largest := i*2+1, i*2+2, i
	//如果左子树存在且比根节点大，则将最大子树下标设置为左子树下标
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	// 如果右子树存在且比根节点大，则将最大子树下标设置为右子树下标
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	// 如果最大子树下标不等于根节点下标，则交换根节点和最大子树下标对应的元素
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		// 递归调用
		maxHeapify(a, largest, heapSize)
	}
}
