package t1146

import "sort"

// SnapshotArray 1146. 快照数组
// https://leetcode.cn/problems/snapshot-array
type pair struct{ snapId, val int }

type SnapshotArray struct {
	curSnapId int
	history   map[int][]pair // 每个 index 的历史修改记录
}

func Constructor(length int) SnapshotArray {
	// 忽略 length
	_ = length
	return SnapshotArray{history: map[int][]pair{}}
}

func (sa *SnapshotArray) Set(index, val int) {
	sa.history[index] = append(sa.history[index], pair{sa.curSnapId, val})
}

func (sa *SnapshotArray) Snap() int {
	sa.curSnapId++
	return sa.curSnapId - 1
}

func (sa *SnapshotArray) Get(index, snapId int) int {
	h := sa.history[index]
	// 找快照编号 <= snapId 的最后一次修改记录
	// 等价于找快照编号 >= snapId+1 的第一个修改记录，它的上一个就是答案
	j := sort.Search(len(h), func(j int) bool { return h[j].snapId >= snapId+1 }) - 1
	if j >= 0 {
		return h[j].val
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
