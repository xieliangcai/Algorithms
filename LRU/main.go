package main

import "fmt"

type LinkNode struct {
	Pre   *LinkNode
	Next  *LinkNode
	key   int
	value int
}

type LRUCache struct {
	Cache    map[int]*LinkNode
	size     int
	capacity int
	head     *LinkNode
	tail     *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{}
	tail := &LinkNode{Pre: head}
	head.Next = tail
	cache := map[int]*LinkNode{}


	return LRUCache{head: head, tail: tail, capacity: capacity, Cache: cache}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Cache[key]; ok {
		v.Next.Pre = v.Pre
		v.Pre.Next = v.Next
		this.tail.Pre.Next = v
		v.Pre = this.tail.Pre
		this.tail.Pre = v
		v.Next = this.tail
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v , ok := this.Cache[key];ok{
		v.value = value
		v.Next.Pre = v.Pre
		v.Pre.Next = v.Next
		this.tail.Pre.Next = v
		v.Pre = this.tail.Pre
		this.tail.Pre = v
		v.Next = this.tail
	}else {
		node := &LinkNode{key: key, value: value}
		if this.size >= this.capacity{
			// 删除掉一个
			delete(this.Cache, this.head.Next.key)
			this.head.Next.Next.Pre = this.head
			this.head.Next = this.head.Next.Next
		}
		this.Cache[key] = node
		this.tail.Pre.Next = node
		node.Pre = this.tail.Pre
		this.tail.Pre = node
		node.Next = this.tail
		this.size++
	}
}



func main() {
	cache := Constructor(2)
	cache.Put(1,1)
	cache.Put(2,2)
	fmt.Println(cache.Get(1))
	cache.Put(3,3)
	fmt.Println(cache.Get(2))
	cache.Put(4,4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

