package t2303

// https://leetcode.cn/problems/calculate-amount-paid-in-taxes/
func calculateTax(brackets [][]int, income int) float64 {
	if income <= 0 {
		return 0
	}

	var tax float64
	for i := 0; i < len(brackets); i++ {
		if income >= brackets[i][0] {
			if i > 0 {
				tax += float64((brackets[i][0] - brackets[i-1][0]) * brackets[i][1])
			} else {
				tax += float64(brackets[i][0] * brackets[i][1])
			}
		} else {
			if i > 0 {
				tax += float64((income - brackets[i-1][0]) * brackets[i][1])
				return tax / 100
			} else {
				tax += float64(income * brackets[i][1])
				return tax / 100
			}
		}
	}
	return tax / 100
}

// https://leetcode.cn/problems/calculate-amount-paid-in-taxes/
func calculateTax1(brackets [][]int, income int) float64 {
	var ans, prev int
	for _, e := range brackets {
		upper, percent := e[0], e[1]
		ans += max(0, min(income, upper)-prev) * percent
		prev = upper
	}
	return float64(ans) / 100.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
