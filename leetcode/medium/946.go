package medium

func validateStackSequences(pushed []int, popped []int) bool {
	st := []int{}
	j := 0
	for _, x := range pushed {
		//将数组中的每一个元素入栈
		st = append(st, x)
		//每次将pushed 的元素入栈之后，如果栈不为空且栈顶元素与 popped 的当前元素相同，
		//则将栈顶元素出栈，同时遍历数组 popped，直到栈为空或栈顶元素与 popped 的当前元素不同。
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}
