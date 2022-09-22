package easy

func canFormArray(arr []int, pieces [][]int) bool {
	index := make(map[int]int, len(pieces))

	//使用哈希表记录下pieces中每个数组的首元素和数组下标的关系
	for i, p := range pieces {
		index[p[0]] = i
	}

	for i := 0; i < len(arr); {
		j, ok := index[arr[i]]
		//如果当前元素不位于哈希表中，则直接返回false
		if !ok {
			return false
		}
		//找到对应的pieces[j] 遍历
		for _, x := range pieces[j] {
			//如果pieces[j]中的元素与arr[i]及以后的元素不相同。则返回false
			if arr[i] != x {
				return false
			}
			i++
		}
	}
	return true
}
