package t705

import "container/list"

const base = 769 // 选择一个质数作为数组大小

// MyHashSet 705.设计哈希集合
// https://leetcode.cn/problems/design-hashset
type MyHashSet struct {
	data []list.List // 双向链表数组 链地址法构造 HashMap
}

func Constructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (s *MyHashSet) hash(key int) int {
	return key % base
}

func (s *MyHashSet) Add(key int) {
	if !s.Contains(key) { // 如果 key 不存在 则添加到计算出来的哈希值的链表后
		h := s.hash(key)
		s.data[h].PushBack(key)
	}
}

func (s *MyHashSet) Remove(key int) {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			s.data[h].Remove(e)
		}
	}
}

// Contains 根据 key 找到对应的链表，遍历链表，返回 key 是否存在
func (s *MyHashSet) Contains(key int) bool {
	h := s.hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
