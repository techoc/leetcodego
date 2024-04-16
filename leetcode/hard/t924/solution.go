package t924

// 924. 尽量减少恶意软件的传播
// https://leetcode.cn/problems/minimize-malware-spread/
// minMalwareSpread 计算最小恶意软件传播范围
// graph：表示网络连接的二维数组，graph[i][j]为1表示节点i和节点j之间有连接，为0表示没有连接；
// initial：表示初始感染的节点数组；
// 返回值：返回移除一个节点后，最小的恶意软件传播范围的节点编号。
func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)               // 节点数量
	ids := make([]int, n)         // 用于标记节点所属的连通分量
	idToSize := make(map[int]int) // 记录每个连通分量的大小
	id := 0                       // 连通分量的唯一标识
	// 通过DFS遍历图，将节点划分为连通分量，并记录每个连通分量的大小
	for i := range ids {
		if ids[i] == 0 {
			id++
			ids[i] = id
			size := 1
			q := []int{i}
			for len(q) > 0 {
				u := q[0]
				q = q[1:]
				for v := range graph[u] {
					if ids[v] == 0 && graph[u][v] == 1 {
						size++
						q = append(q, v)
						ids[v] = id
					}
				}
			}
			idToSize[id] = size
		}
	}
	idToInitials := make(map[int]int) // 记录初始感染的连通分量数量
	for _, u := range initial {
		idToInitials[ids[u]]++
	}
	ans := n + 1    // 初始化答案为一个不可能达到的值
	ansRemoved := 0 // 初始化移除的节点数量为0
	// 遍历初始感染的节点，计算并更新最小恶意软件传播范围的节点
	for _, u := range initial {
		removed := 0
		if idToInitials[ids[u]] == 1 {
			removed = idToSize[ids[u]] // 如果连通分量中只有一个初始感染节点，则移除该节点后，整个连通分量都被移除
		}
		if removed > ansRemoved || (removed == ansRemoved && u < ans) {
			ans, ansRemoved = u, removed
		}
	}
	return ans // 返回最小恶意软件传播范围的节点编号
}
