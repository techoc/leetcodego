package medium

func groupThePeople(groupSizes []int) (ans [][]int) {
	groups := map[int][]int{}
	//遍历group
	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}
	for size, people := range groups {
		for i := 0; i < len(people); i += size {
			ans = append(ans, people[i:i+size])
		}
	}
	return
}
