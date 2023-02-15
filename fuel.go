package leetcode

import "math"

func minimumFuelCost(roads [][]int, seats int) int64 {
	gr := make([][]int, len(roads)+1)
	for _, r := range roads {
		gr[r[0]] = append(gr[r[0]], r[1])
		gr[r[1]] = append(gr[r[1]], r[0])
	}

	var fuel int64
	dfs(-1, 0, gr, int64(seats), &fuel)
	return fuel
}

func dfs(prev int, cur int, gr [][]int, seats int64, fuel *int64) int64 {
	var riders int64 = 1

	for _, next := range gr[cur] {
		if next != prev {
			riders += dfs(cur, next, gr, seats, fuel)
		}
	}

	if cur != 0 {
		*fuel += int64(math.Ceil(float64(riders) / float64(seats)))
	}

	return riders
}
