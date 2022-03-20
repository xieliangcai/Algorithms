package main

import "fmt"

// 归并排序

func mergeSort(data []int) {
	mergerSort_C(data, 0, len(data))
}

func mergerSort_C(data []int, start, end int) {
	if end-start <= 1 {
		return
	}
	mid := (start + end) / 2
	mergerSort_C(data, start, mid)
	mergerSort_C(data, mid, end)
	merge(data, start, end, mid)
	fmt.Println(data)
}

// 合并两个数组
func merge(data []int, start, end, mid int) {
	listA := data[start:mid]
	listB := data[mid:end]
	fmt.Println(listA, listB)
	temp := []int{}
	indexA := start
	indexB := mid
	for indexA < mid && indexB < end {
		if data[indexA] <= data[indexB] {
			temp = append(temp, data[indexA])
			indexA++
		} else {
			temp = append(temp, data[indexB])
			indexB++
		}
	}

	if indexA < mid {
		temp = append(temp, data[indexA:mid]...)
	}

	if indexB < end {
		temp = append(temp, data[indexB:end]...)
	}

	for i := 0; i < end-start && i < len(temp); i++ {
		data[start+i] = temp[i]
	}
	return

}

func main() {
	list := []int{1,2,-1,-1}
	mergeSort(list)
	fmt.Println(list)
}
