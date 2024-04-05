package t2798

// 2798.满足目标工作时长的员工数目
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
