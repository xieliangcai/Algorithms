package main

import "fmt"

// 插入排序
/*
1、在已排序区间中找到合适的插入位置将其插入, 将后面的都往后移动一位， 并保持已排序区间数据一直有序
难点：找到插入点
*/

func InsertSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	for i := 1; i < length; i++ {
		value := data[i]
		for j := i-1; j >= 0; j-- {
			if data[j] > value {
				data[j+1] = data[j]
			}else {
				break
			}
			data[j] = value
		}
	}
	return data

}

// data 是一个有序表,查找插入位置
func find(data []int, findData int) (index int) {
	for k, v := range data {
		if findData < v {
			return k
		}
	}
	return 0
}

// 往后移动，然后插入
func move(dataList []int, data int, index int, length int) []int {
	for i := length; i >= index; i-- {
		dataList[i+1] = dataList[i]
	}
	dataList[index] = data
	return dataList

}

func main() {
	list := []int{33, 3, 6, 5, 11, 1, 4}
	//result := InsertSort(list)
	//fmt.Println(result)
	a(list)
	dmap := map[string]int{"a":1}
	b(dmap)
	fmt.Println(dmap)

}

func a(data []int)  {
	data[0] = 3
	return

}

func b(data map[string]int)  {
	data["a"] = 2
	return

}
