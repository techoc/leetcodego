package t1822

func arraySign(nums []int) int {
	//初始化元素为负数的个数
	count := 0
	for _, num := range nums {
		//只要有一个元素为0 就直接返回0
		if num == 0 {
			return 0
		}
		if num < 0 {
			count += 1
		}
	}
	//如果count的个数为奇数
	//就返回 -1
	if count%2 != 0 {
		return -1
	}
	return 1
}
