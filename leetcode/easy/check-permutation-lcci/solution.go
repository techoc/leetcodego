package check_permutation_lcci

func CheckPermutation(s1 string, s2 string) bool {
	//维护两个hashmap
	var c1, c2 [128]int
	for _, ch := range s1 {
		c1[ch]++
	}
	for _, ch := range s2 {
		c2[ch]++
	}
	return c1 == c2
}
