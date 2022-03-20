package main

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var tail *ListNode
	// 上次的进位
	last := 0
	for ;l1 != nil || l2 != nil; {
		var v1, v2 int
		if l1 != nil{
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			v2 =  l2.Val
			l2 = l2.Next
		}
		value := v1 + v2 + last
		last = value/10
		sum := value%10
		if result == nil{
			result = &ListNode{Val: sum}
			tail = result
			continue
		}
		tail.Next = &ListNode{sum, nil}
		tail = result.Next
	}

	if last != 0{
		tail.Next =  &ListNode{last, nil}
	}
	return result
}



func addNode(a *ListNode, v int){
	for ; a.Next != nil; {
		a = a.Next
	}
	a.Next = &ListNode{v, nil}
	return
}

func main() {
	l1 := &ListNode{2, nil}
	l2 := &ListNode{5, nil}
	addNode(l1, 4)
	addNode(l2, 6)
	addNode(l1, 3)
	addNode(l2, 4)
	addTwoNumbers(l1, l2)
}