package t1185

// 1185.一周中的第几天
// https://leetcode.cn/problems/day-of-the-week
func dayOfTheWeek(day int, month int, year int) string {
	weeks := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// 蔡勒公式
	if month <= 2 {
		month += 12
		year--
	}
	c := year / 100
	y := year % 100
	w := (c/4 - 2*c + y + y/4 + 13*(month+1)/5 + day - 1) % 7
	// week 计算出来可能是负数 则加上7 变成正数 后取模
	return weeks[(w+7)%7]
}
