package t2894

// 2894.分类求和并做差
// https://leetcode.cn/problems/divisible-and-non-divisible-sums-difference/
func differenceOfSums(n int, m int) int {
	num1 := 0
	num2 := 0
	for i := 1; i <= n; i++ {
		if i%m != 0 { // 如果不能被m整除，则加入num1
			num1 += i
		} else { // 如果能被m整除，则加入num2
			num2 += i
		}
	}
	return num1 - num2
}
