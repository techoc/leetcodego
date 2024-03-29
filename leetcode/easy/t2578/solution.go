package t2578

// https://leetcode.cn/problems/split-with-minimum-sum/
func splitNum(num int) int {
	//新建一个map 存储 0~9 在num出现的次数
	cnt := [10]int{}
	//记录 nums 的长度
	n := 0
	for ; num > 0; num /= 10 {
		cnt[num%10]++
		n++
	}

	ans := [2]int{}
	for i, j := 0, 0; i < n; i++ {
		// 如果当前数字出现次数为0 则跳过
		for cnt[j] == 0 {
			j++
		}
		// 将当前数字加入结果集
		cnt[j]--
		ans[i&1] = ans[i&1]*10 + j
	}

	return ans[0] + ans[1]
}
