package hard

func uniqueLetterString(s string) (ans int) {
	dict := map[rune][]int{}
	for i, c := range s {
		if v, ok := dict[c]; ok {
			dict[c] = append(v, i)
		} else {
			dict[c] = []int{-1, i}
		}
	}
	for _, v := range dict {
		v = append(v, len(s))
		for i := 1; i < len(v)-1; i++ {
			ans += (v[i] - v[i-1]) * (v[i+1] - v[i])
		}
	}
	return
}
