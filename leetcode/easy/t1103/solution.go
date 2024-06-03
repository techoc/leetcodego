package t1103

// 1103. 分糖果 II
// https://leetcode.cn/problems/distribute-candies-to-people
func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	for i := 0; candies > 0; i++ {
		ans[i%num_people] += min(candies, i+1)
		candies -= i + 1
	}
	return ans
}
