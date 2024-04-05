package t1137

// 1137.第N个泰波那契数
// https://leetcode.cn/problems/n-th-tribonacci-number/
func tribonacci(n int) int {
	// 当 n 为 0 时，直接返回 0
	if n == 0 {
		return 0
	}
	// 当 n 小于等于 2 时，直接返回 1
	if n <= 2 {
		return 1
	}
	// 初始化变量，p, q, r, s 分别为 tribonacci 数的前四个值
	p, q, r, s := 0, 0, 1, 1
	// 从第 3 个数开始计算，循环直到计算到第 n 个数
	for i := 3; i <= n; i++ {
		// 更新变量值，每次循环都把当前的 p, q, r 的值更新为 q, r, s
		p = q
		q = r
		r = s
		// 计算下一个 tribonacci 数的值
		s = p + q + r
	}
	// 循环结束后，返回第 n 个 tribonacci 数的值
	return s
}

// 递归 会超时
//func tribonacci(n int) (res int) {
//	// 边界判定
//	if n == 0 {
//		return 0
//	}
//	if n == 1 {
//		return 1
//	}
//	if n == 2 {
//		return 1
//	}
//	res += tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
//	return res
//}
