package t2485

func pivotInteger(n int) int {
	// 从后往前 枚举
	for i := n; i > 0; i-- {
		var sum1, sum2 int

		// 如果 1 和 i 之间的所有元素之和等于 i 和 n 之间所有元素之和。
		for j := 0; j <= i; j++ {
			sum1 += j
		}

		for k := i; k <= n; k++ {
			sum2 += k
		}

		if sum1 == sum2 {
			return i
		}
	}

	return -1
}
