package main

import "fmt"

// 二分查找

func main() {
	row := 5
	col	:= 5
	visited := [][]bool{}
	for i:= 0; i< row;i++{
		visited = append(visited, make([]bool, col))

	}
	fmt.Println(visited)
}

func search(nums []int, target int) int {
	low := 0
	hight := len(nums) - 1
	for low <= hight {
		mid := low + (hight-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func hasPath(maze [][]int, start []int, destination []int) bool {
	row, col := len(maze), len(maze[0])

	visited := [][]bool{}
	for i:= 0; i< row;i++{
		visited = append(visited, make([]bool, col))

	}


	var dfs func(maze [][]int, start []int, destination []int,visited  [][]bool)bool
	dfs = func(maze [][]int, start []int,destination []int, visited [][]bool)bool{
		if visited[start[0]][start[1]]{
	return false
	}
	if start[0] == destination[0] && start[1] == destination[1]{
	return true
	}

	visited[start[0]][start[1]] = true

	up := start[0] -1;
	down := start[0] + 1;
	left := start[1] -1;
	right := start[1]+1;

	for up >= 0 && maze[up][start[1]] == 0{
	up--
	}
	if (dfs(maze, []int{up+1, start[1]}, destination, visited)){
	return true
	}

	for down <= row && maze[down][start[1]] == 0{
	down++
	}
	if (dfs(maze, []int{down-1, start[1]}, destination, visited)){
	return true
	}

	for left >=0&& maze[start[0]][left]==0{
	left--
	}
	if dfs(maze, []int{start[0], left+1}, destination, visited){
	return true
	}

	for right < col && maze[start[0]][right]==0{
	right++
	}
	if dfs(maze, []int{start[0], right-1}, destination, visited){
	return true
	}

	return false
	}
	return dfs(maze, start, destination, visited)
}
