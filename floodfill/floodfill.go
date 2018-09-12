package main

type pos struct {
	row    int
	column int
}

//maxConnectedValue returns the value and the connection count of the grid element
//with most connections in the given grid.
//Grid elementes are connected when they are adjacent and they have the same value.
func maxConnectedValue(grid [][]int) (int, int) {
	visited := make(map[pos]bool)
	maxValue := -1
	maxCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			me := pos{i, j}
			count := dfs(me, visited, grid)
			if count > maxCount {
				maxCount = count
				maxValue = grid[me.row][me.column]
			}
		}
	}

	return maxValue, maxCount
}

func dfs(c pos, visited map[pos]bool, grid [][]int) int {
	if visited[c] {
		return 0
	}
	visited[c] = true
	count := 1
	neighbours := unvisitedNeighbours(c, visited, grid)
	for _, neighbour := range neighbours {
		count += dfs(neighbour, visited, grid)
	}
	return count
}

func unvisitedNeighbours(c pos, visited map[pos]bool, grid [][]int) []pos {
	x, y := c.row, c.column
	value := grid[x][y]

	var result []pos

	if x-1 >= 0 && !visited[pos{x - 1, y}] && grid[x-1][y] == value {
		result = append(result, pos{x - 1, y})
	}
	if y-1 >= 0 && !visited[pos{x, y - 1}] && grid[x][y-1] == value {
		result = append(result, pos{x, y - 1})
	}
	if x+1 < len(grid) && !visited[pos{x + 1, y}] && grid[x+1][y] == value {
		result = append(result, pos{x + 1, y})
	}
	if y+1 < len(grid[x]) && !visited[pos{x, y + 1}] && grid[x][y+1] == value {
		result = append(result, pos{x, y + 1})
	}

	return result
}
