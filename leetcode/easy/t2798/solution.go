package t2798

// 2798. 统计满足要求三元组的数目
// https://leetcode.cn/problems/number-of-employees-who-met-the-target/
func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	res := 0
	for _, h := range hours {
		if h >= target {
			res++
		}
	}
	return res
}
