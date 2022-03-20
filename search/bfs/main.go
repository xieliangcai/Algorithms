package main

import "fmt"

/*
岛屿数量-200
https://leetcode-cn.com/problems/number-of-islands/
*/

func main() {
	grid := [][]string{}
	grid = append(grid, []string{"1", "1", "1"})
	grid = append(grid, []string{"0", "1", "0"})
	grid = append(grid, []string{"1", "1", "1"})
	//grid = append(grid, []string{"0", "0", "0", "0", "0"})
	result := numIslands(grid)
	fmt.Println(result)

}

func numIslands(grid [][]string) int {
	row := len(grid)
	col := len(grid[0])
	// 标识4个位置
	qx := []int{-1, 0, 1, 0}
	qy := []int{0, -1, 0, 1}

	visited := [][]bool{}
	for i:=0; i< row ; i++{
		visited = append(visited, make([]bool, col))
	}
	var bfs func(x, y int)
	bfs = func(x, y int) {
		if x>=row || y>=col{
			return
		}
		// 先进先出
		// 找到临近块为1的
		for i := 0; i < 4; i++ {
			nx := x + qx[i]
			ny := y + qy[i]
			if nx < 0 || nx >= row || ny < 0 || ny >= col {
				continue
			}

			if grid[nx][ny] == "1"&& !visited[nx][ny]{
				visited[nx][ny] = true
				bfs(nx, ny)
			}
		}
	}

	var ans int
	for rowIndex, row := range grid {
		for colIndex, value := range row {
			if value == "1" && !visited[rowIndex][colIndex] {
				ans++
				bfs(rowIndex, colIndex)
			}
		}
	}
	return ans
}

