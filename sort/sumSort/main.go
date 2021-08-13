package main

import "fmt"

// 计数排序（桶排序的升级）2 5 3 0 2 3 0 3
// 汇总每个阶段的人数  {0:2, 1:0, 2:2, 3:3, 4:0, 5:1}
// 通过人数累计每个阶段的位置  {0:2, 1:2, 2:4, 3:7, 4:0, 5:8}
// 开始排序， 现在保存的是每个阶段的最大位置， 每次遍历减一， 得到真实的位置如3的位置是 第5、6、7 位

func sumSort(data []int, maxValue int) []int  {
	countMap := map[int]int{}
	// 初始每个阶段的人数
	for i:=0; i<=maxValue; i++ {
		countMap[i] = 0
	}

	// 统计每个阶段的人数
	for _, v := range data{
		countMap[v]++
	}

	var totalCount int

	// 转换成么个阶段人数的位置（核心逻辑）
	for i:=0; i<=maxValue; i++ {
		totalCount+=countMap[i]
		countMap[i] = totalCount
	}

	newData := make([]int, totalCount)
	// 排序
	for _, v := range data{
		countMap[v]--
		newData[countMap[v]] = v
	}
	return newData
}

func main() {
	dataList := []int{2,5,3,0,2,3,0,3}
	dataList = sumSort(dataList, 5)
	fmt.Println(dataList)
}