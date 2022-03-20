package main

import "fmt"

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
url： https://leetcode-cn.com/problems/container-with-most-water/
*/

func main()  {
	a := []int{1,8,6,2,5,4,8,3,7}
	result := maxArea(a)
	fmt.Println(result)
}

func maxArea(height []int) int {
	// 第一格是空的
	length := len(height)-1
	if length < 1{
		return 0
	}

	var maxArea int
	end:=length
	// 双指针
	for start:= 0; start<length; {
		if start == end{
			break
		}
		fmt.Println(start, end)
		if maxArea < (end-start)*minValue(height[start], height[end]){
			maxArea = (end-start)*minValue(height[start], height[end])
		}
		// 比两数大小， 往小的一侧移动
		if height[start] <= height[end]{
			start++
		}else{
			end--
		}
	}
	return maxArea

}

func minValue(x, y int) int{
	if x<y{
		return  x
	}
	return y
}