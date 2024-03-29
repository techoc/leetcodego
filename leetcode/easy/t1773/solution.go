package t1773

var d = map[string]int{"type": 0, "color": 1, "name": 2}

func countMatches(items [][]string, ruleKey, ruleValue string) (ans int) {
	index := d[ruleKey]
	for _, item := range items {
		if item[index] == ruleValue {
			ans++
		}
	}
	return
}
