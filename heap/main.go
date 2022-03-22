package main

import "fmt"

func main()  {
	nums :=[]int{1,3,-1,-3,5,3,6,7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}




func maxSlidingWindow(nums []int, k int) []int {
	maxList  := []int{}
	q := []int{}
	// 保持最大数在q[0]位置
	push := func(i int){
		for len(q)>0  && nums[i] >= nums[q[len(q)-1]]{
			q = q[:len(q)-1]
		}

		q = append(q, i)
		return
	}
	for i:= 0;i< k; i++{
		push(i)
	}
	maxList = append(maxList, nums[q[0]])
	for i := k; i< len(nums); i++{
		push(i)
		// 删除不在窗口的数据
		for q[0]<=i-k{
			q = q[1:]
		}
		maxList = append(maxList, nums[q[0]])
	}

	return maxList
}




