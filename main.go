package main

import (
	"fmt"
	"sort"
)

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
	coins := []int{411,412,413,414,415,416,417,418,419,420,421,422}
	amount := 9864
	result := coinChange(coins, amount)
	fmt.Println(result)
}

func coinChange(coins []int, amount int) int {
	if amount == 0{
		return 0
	}
	sort.Ints(coins)
	ans := int(^uint(0)>>1)
	var change func(coins []int,amount, c_index, count, a int)
	change = func(coins []int,amount, c_index, count, a int){
		if (amount==0){
			a = min(a, count)
			ans = min(a, ans)
			return
		}
		if (c_index == len(coins)){
			return
		}
		for k :=(amount/coins[c_index]); k>=0 && k+count< a; k--{
			change(coins, amount-k*coins[c_index], c_index+1, count+k, a)
		}
	}
	change(coins, amount, 0,0, ans)
	if ans == int(^uint(0)>>1){
		return -1
	}
	return ans

}

func min( a, b int)int{
	if a>=b{
		return b
	}
	return a
}