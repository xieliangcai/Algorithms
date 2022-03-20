package hashMap

import (
	"Algorithms/hashTable/linkNodes"
)

type HashTable struct {
	arr [16]*linkNodes.NodeList
}

func (this *HashTable) Add(k, v string) {
	index := hashCode(k)
	if len(this.arr) == 0{
		this.arr = [16]*linkNodes.NodeList{}
	}
	if this.arr[index] == nil{
		this.arr[index] = &linkNodes.NodeList{Header: &linkNodes.Node{Data: linkNodes.MP{K: k,V: v}, Next: nil}}
		return
	}
	this.arr[index].Add(k, v)
}

func (this *HashTable) Search(k string)(v string, ok bool) {
	if len(this.arr) == 0{
		return
	}
	index := hashCode(k)
	if this.arr[index] == nil{
		return
	}
	return this.arr[index].Search(k)
}

// 计算下标索引
func hashCode(key string) int {
	var index int
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		index = index * (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1
	return index
}
