package t706

import "container/list"

type entry struct {
	key   int
	value int
}

var base = 769

// MyHashMap 706. 设计哈希映射
// https://leetcode.cn/problems/design-hashmap
type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{
		data: make([]list.List, base),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	// 遍历链表 找到key 如果找到就替换值
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			e.Value = entry{key, value}
			return
		}
	}
	// 否则插入
	this.data[h].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	// 遍历链表 找到key 如果找到就返回值
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			return e.Value.(entry).value
		}
	}
	// 否则返回 -1
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	// 遍历链表 找到key 如果找到就删除
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			this.data[h].Remove(e)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
