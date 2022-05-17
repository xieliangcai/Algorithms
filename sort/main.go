package main

import (
	"fmt"
)

func sort(data []int) []int {
	//result := bubbleSort.BubbleSort(data)
	result := InsertSort(data)
	fmt.Println(result)
	return result
}


func InsertSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}

	for i := 1; i<length; i++{
		value := data[i]
		j := i -1
		// 找到插入點
		for ;j>=0; j--{
			if data[j] > value{
				data[j+1] = data[j]
			}else {
				break
			}

			data[j] = value

		}

	}


	//
	//for i := 0; i < len(data)-1; i++ {
	//	if data[i] <= data[i+1] {
	//		continue
	//	}
	//
	//	if i == 0 {
	//		temp := data[i]
	//		data[i] = data[i+1]
	//		data[i+1] = temp
	//		continue
	//
	//	}
	//
	//	for j := 0; j <= i; j++ {
	//
	//
	//		if data[i] > data[j] && data[i] < data[j+1] {
	//			temp := data[i]
	//			//往后移动一位
	//			for d := i; d >= j+1; d-- {
	//				data[d] = data[d-1]
	//			}
	//
	//
	//			data[j+1] = temp
	//
	//		}
	//
	//	}
	//}
	return data

}


func check(result, expectList []int) bool {
	if len(result) != len(expectList) {
		return false
	}
	for k, v := range result {
		if v != expectList[k] {
			return false
		}
	}
	return true
}

func main() {
	list := []int{33, 3, 6, 5, 11, 1, 4}
	expectList := []int{1, 3, 4, 5, 6, 11, 33}


	result := sortArray(list)
	fmt.Println(check(result, expectList))

}

func sortArray(nums []int) []int {
	// 选择排序
	// 找到最小值，放到最前面
	for i, v := range nums {
		min := v
		fmt.Printf("%v \n",min)
		var minIndex int
		for j:= i+1; j< len(nums); j++{
			fmt.Printf("j %v %v \n", j, min)
			if nums[j] < min{
				min = nums[j]
				minIndex = j
			}
		}
		fmt.Printf("min %v %v \n", min, nums[minIndex])
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}