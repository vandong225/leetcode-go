package graphgeneral

func numIslands(grid [][]byte) int {
	var fillWater func(i, j int)

	fillWater = func(i, j int) {

		if j >= len(grid[0]) || i < 0 || j < 0 || i >= len(grid) || grid[i][j] != '1' {
			return
		}

		grid[i][j] = '0'
		fillWater(i+1, j)
		// fillWater(i+1, j+1)
		// fillWater(i+1, j-1)
		fillWater(i, j+1)
		fillWater(i, j-1)
		fillWater(i-1, j)
		// fillWater(i-1, j+1)
		// fillWater(i-1, j-1)
	}

	nums := 0

	for i, r := range grid {
		for j, c := range r {
			if c == '1' {
				nums++
				fillWater(i, j)
			}
		}
	}

	return nums
}
