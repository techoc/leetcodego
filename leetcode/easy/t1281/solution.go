package t1281

// 1281. 整数的各位积和之差
// https://leetcode.cn/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
func subtractProductAndSum(n int) int {
	product, sum := 1, 0
	for n > 0 { // 取得位数
		product *= n % 10
		sum += n % 10
		n /= 10
	}
	return product - sum
}
