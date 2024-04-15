package t100256

import "fmt"

// 100256.替换字符可以得到的最短时间
// https://leetcode.cn/problems/latest-time-you-can-obtain-after-replacing-characters
func findLatestTime(s string) string {
	// 倒叙枚举
	for i := 11; i >= 0; i-- {
		for j := 59; j >= 0; j-- {
			ok := true
			time := fmt.Sprintf("%02d:%02d", i, j) // 枚举出时间
			for index, value := range s {          // 如果 不等于 ? 且 不相等
				if value != '?' && time[index] != byte(value) { // 则不满足条件
					ok = false
				}
			}
			if ok {
				return time
			}
		}
	}
	return s
}
