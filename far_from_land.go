package leetcode

func maxDistance(grid [][]int) int {
	dist := 0
	n := len(grid)
	x := []int{-1, 0, 1, 0}
	y := []int{0, -1, 0, 1}
	q := make([][2]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	if len(q) == n*n {
		return -1
	}

	for len(q) > 0 {
		dist++

		for qlen := len(q); qlen > 0; qlen-- {
			p := q[0]
			q = q[1:]

			for dir := 0; dir < 4; dir++ {
				i := p[0] + x[dir]
				j := p[1] + y[dir]

				if i >= 0 && i < n && j >= 0 && j < n && grid[i][j] == 0 {
					q = append(q, [2]int{i, j})
					grid[i][j] = 1
				}
			}
		}
	}

	return dist - 1
}
