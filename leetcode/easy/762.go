package easy

func countPrimeSetBits(left int, right int) (ans int) {
	//一个 int 的二进制表示不超过 32,列出32以内的所有素数
	hash := map[int]bool{
		2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true,
	}

	for i := left; i <= right; i++ {
		cur := 0
		for j := i; j > 0; j -= j & (-j) {
			cur++
		}
		if hash[cur] {
			ans++
		}
	}
	return
}
