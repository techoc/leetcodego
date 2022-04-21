package easy

func majorityElement(nums []int) int {
	m := make(map[int]int)
	// 初始化map, 将所有元素放入map中
	for _, v := range nums {
		m[v]++
	}
	// 遍历map, 如果map中的值大于n/2, 则返回该值
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
