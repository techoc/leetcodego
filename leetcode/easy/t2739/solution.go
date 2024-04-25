package t2739

// 2739. 总行驶距离
// https://leetcode.cn/problems/total-distance-traveled
func distanceTraveled(mainTank int, additionalTank int) int {
	ans := 0
	for mainTank >= 5 { // 如果主油箱有 5 升油，则消耗 50 升油
		ans += 50
		mainTank -= 5
		if additionalTank > 0 { // 如果副油箱有油，则补充 1 升油
			mainTank++
			additionalTank--
		}
	}
	ans += mainTank * 10
	return ans
}

func distanceTraveledTwo(mainTank, additionalTank int) int {
	res := 0 // 初始化总行进距离为0

	// 当主油箱中的油量大于等于5时，进行循环计算
	for mainTank >= 5 {
		// 每次至少能行进5单位距离，累加到结果中
		res += mainTank / 5 * 5
		// 计算本次能从主油箱中取出的油量，但不超过附加油箱的剩余油量
		addNum := (mainTank / 5)
		if addNum > additionalTank {
			addNum = additionalTank
		}
		// 更新主油箱和附加油箱的油量
		mainTank = mainTank%5 + addNum // 主油箱剩余油量加上取出的油量
		additionalTank -= addNum       // 附加油箱减去取出的油量
	}

	// 最后，将剩余的油量转换为行进距离，并累加到结果中，乘以10表示单位转换
	return (res + mainTank) * 10
}
