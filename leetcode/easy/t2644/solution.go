package t2644

// 2644. 找出可整除性得分最大的整数
// https://leetcode.cn/problems/find-the-maximum-divisibility-score
func maxDivScore(nums []int, divisors []int) int {
	res, res_count := divisors[0], 0
	for _, value := range divisors {
		count := 0
		for _, num := range nums {
			if num%value == 0 { // 整除
				count++
			}
		}
		if count > res_count || (count == res_count && value < res) {
			res = value
			res_count = count
		}
	}
	return res
}
