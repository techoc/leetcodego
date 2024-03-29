package easy

func maximumWealth(accounts [][]int) int {
	var max int
	for _, v := range accounts {
		var sum int
		for _, vv := range v {
			sum += vv
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
