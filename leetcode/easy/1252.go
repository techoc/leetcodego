package easy

func oddCells(m, n int, indices [][]int) (ans int) {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for _, p := range indices {
		for j := range matrix[p[0]] {
			matrix[p[0]][j]++
		}
		for _, row := range matrix {
			row[p[1]]++
		}
	}
	for _, row := range matrix {
		for _, v := range row {
			ans += v % 2
		}
	}
	return
}
