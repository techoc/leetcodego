package medium

const one int = 1 << 7
const two int = one + (1 << 6)

func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		l := 1
		for l < 7 && (data[i]>>(8-l))&1 == 1 {
			l++
		}
		if l == 2 || l > 5 || i+l-2 >= len(data) {
			return false
		}
		if l > 2 {
			l--
		}
		for j := i + 1; j < i+l; j++ {
			if data[j]&two != one {
				return false
			}
		}
		i += l
	}
	return true
}
