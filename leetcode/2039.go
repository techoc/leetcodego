package leetcode

import "math"

func networkBecomesIdle(edges [][]int, patience []int) (ret int) {
	//建图
	//BFS从主节点获取从节点到主节点的最短路径
	//计算最后一次重新发送的时间，以及之后需要多久到达服务器

	visited, cache := make([]bool, 100001), make([][]int, len(patience))
	visited[0] = true
	for i := 0; i < len(patience); i++ {
		cache[i] = []int{}
	}
	for _, edge := range edges {
		cache[edge[0]] = append(cache[edge[0]], edge[1])
		cache[edge[1]] = append(cache[edge[1]], edge[0])
	}

	var step []int
	step = append(step, 0)
	for i := 1; len(step) != 0; i++ {
		var level []int
		for _, j := range step {
			for _, k := range cache[j] {
				if !visited[k] {
					visited[k], ret = true, int(math.Max(float64(ret), float64(i*2+(i*2-1)/patience[k]*patience[k])))
					level = append(level, k)
				}
			}
		}
		step = level
	}
	return ret + 1
}
