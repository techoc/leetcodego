package t1668

func maxRepeating(sequence, word string) (ans int) {
	n, m := len(sequence), len(word)
	if n < m {
		return
	}
	f := make([]int, n)
	for i := m - 1; i < n; i++ {
		if sequence[i-m+1:i+1] == word {
			if i == m-1 {
				f[i] = 1
			} else {
				f[i] = f[i-m] + 1
			}
			if f[i] > ans {
				ans = f[i]
			}
		}
	}
	return
}
