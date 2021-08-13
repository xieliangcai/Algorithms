package main

import "fmt"

type Node struct {
	Data string `json:"data"`
	Next *Node  `json:"next"`
}

type NodeList struct {
	Header *Node `json:"header"`
}

func (this *NodeList) add(data string) {
	if this.length() == 0 {
		this.Header = &Node{Data: data}
		return
	}
	cur := this.Header
	for ; cur.Next != nil; {
		cur = cur.Next
	}
	cur.Next = &Node{Data: data}
}

func (this *NodeList) remove(data string) {
	if this.length() == 0 {
		return
	}

	cur := this.Header
	for ; cur.Next != nil; {
		if cur.Data == data {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
}

func (this *NodeList) search(data string) {
	if this.length() == 0 {
		return
	}

	cur := this.Header
	i := 0
	for ; cur.Next != nil; i++ {
		if cur.Data == data {
			fmt.Println(i)
		}
		cur = cur.Next
	}

	if cur.Data == data{
		fmt.Println(i)
		return
	}

}

func (this *NodeList) traverse() {
	if this.length() == 0 {
		return
	}
	cur := this.Header
	for ; cur.Next != nil; {
		fmt.Println(cur.Data)
		cur = cur.Next
	}
	fmt.Println(cur.Data)
}

func (this *NodeList) length() int {
	if this.Header == nil {
		return 0
	}

	cur := this.Header
	count:=1
	for ; cur.Next != nil; {
		count++
		cur = cur.Next
	}
	return count
}

func main() {
	nodeList := &NodeList{}
	nodeList.add("a")
	nodeList.traverse()
	nodeList.add("b")
	nodeList.traverse()
	nodeList.search("b")
}
