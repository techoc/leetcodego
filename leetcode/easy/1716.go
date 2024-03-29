package easy

// 因为每周七天存的钱之和比上一周多 77 块，因此每周存的钱之和的序列是一个等差数列，我们可以用等差数列求和公式来求出所有完整的周存的钱总和。
//剩下的天数里，每天存的钱也是一个等差数列，可以用相同的公式进行求和。最后把两者相加可以得到答案。
func TotalMoney(n int) (ans int) {
	// 所有完整的周存的钱
	weekNum := n / 7
	firstWeekMoney := (1 + 7) * 7 / 2
	lastWeekMoney := firstWeekMoney + 7*(weekNum-1)
	weekMoney := (firstWeekMoney + lastWeekMoney) * weekNum / 2
	// 剩下的不能构成一个完整的周的天数里存的钱
	dayNum := n % 7
	firstDayMoney := 1 + weekNum
	lastDayMoney := firstDayMoney + dayNum - 1
	dayMoney := (firstDayMoney + lastDayMoney) * dayNum / 2
	return weekMoney + dayMoney
}
