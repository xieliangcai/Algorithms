package linkNodes

import "fmt"

type MP struct {
	K string `json:"k"`
	V string `json:"v"`
}

type Node struct {
	Data MP    `json:"data"`
	Next *Node `json:"next"`
}

type NodeList struct {
	Header *Node `json:"header"`
}

func (this *NodeList) Add(k, v string) {
	if this.Length() == 0 {
		this.Header = &Node{Data: MP{k,  v}}
		return
	}
	cur := this.Header
	for ; cur.Next != nil; {
		cur = cur.Next
	}
	//查找链表中是否已经存在这个k， 如果存在删除旧的，把新的加上
	if _,ok := this.Search(k); ok{
		this.Remove(k)
	}
	cur.Next = &Node{Data: MP{k, v}}
}

func (this *NodeList) Remove(k string) {
	if this.Length() == 0 {
		return
	}
	cur := this.Header
	for ; cur.Next != nil; {
		if cur.Data.K == k {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
}

func (this *NodeList) Search(k string) ( string, bool) {
	if this.Length() == 0 {
		return "", false
	}

	cur := this.Header
	i := 0
	for ; cur.Next != nil; i++ {
		if cur.Data.K == k {
			fmt.Println(i)
			return cur.Data.V, true
		}
		cur = cur.Next
	}

	if cur.Data.K == k {
		fmt.Println(i)
		return cur.Data.V, true
	}

	return "", false
}

func (this *NodeList) Traverse() {
	if this.Length() == 0 {
		return
	}
	cur := this.Header
	for ; cur.Next != nil; {
		fmt.Println(cur.Data.K, cur.Data.V)
		cur = cur.Next
	}
	fmt.Println(cur.Data.K, cur.Data.V)
}

func (this *NodeList) Length() int {
	if this.Header == nil {
		return 0
	}

	cur := this.Header
	count := 1
	for ; cur.Next != nil; {
		count++
		cur = cur.Next
	}
	return count
}
