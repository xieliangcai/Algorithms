package main

import "fmt"

/*
433. 最小基因变化
https://leetcode-cn.com/problems/minimum-genetic-mutation/
*/

func main() {
	start := "AACCGGTT"
	end := "AAACGGTA"
	bank := []string{"AACCGGTA","AACCGCTA","AAACGGTA"}
	fmt.Println(start[0:1])
	result := minMutation(start, end, bank)
	fmt.Println(result)
}

func minMutation(start string, end string, bank []string) int {
	length := len(start)
	ans := -1
	var dfs func(str string, oldIndex, result int, intdex int)

	replaceKey := []string{"A", "C", "G", "T"}

	dfs = func(str string,oldIndex int, result int, index int) {
		if str == end {
			if ans != -1 && result < ans {
				ans = result
				return
			} else {
				ans = result
				return
			}
		}
		if str == "AAACGGTT"{
			fmt.Println(str)
		}
		if index > length {
			return
		}

		// 裁剪 , 仅仅当前缀和bank里相同时，才能往下走
		for _, key := range replaceKey {
			if index<length && key[0] == str[index]{
				continue
			}
			newStr := str[:index-1] + key + str[index:]
			for _, v := range bank {
				if newStr[:index] == v[:index] {
					fmt.Println(newStr)
					dfs(newStr,oldIndex, result+1, index+1)
				}}
		}


	}
	for index := 1; index <= length; index++ {
		dfs(start, index,0, index)
	}
	return ans
}
