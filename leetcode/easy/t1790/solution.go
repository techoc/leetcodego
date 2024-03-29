package t1790

func areAlmostEqual(s1 string, s2 string) bool {
	count := 0
	var c1, c2 byte
	for i := range s1 {
		a, b := s1[i], s2[i]
		if a != b {
			count++
			if count > 2 || (count == 2 && (a != c2 || b != c1)) {
				return false
			}
			c1, c2 = a, b
		}
	}
	return count != 1
}
