package leetcode

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	const red, blue, visited = 0, 1, -1

	gr := make([][][2]int, n)
	for _, e := range redEdges {
		gr[e[0]] = append(gr[e[0]], [2]int{e[1], red})
	}
	for _, e := range blueEdges {
		gr[e[0]] = append(gr[e[0]], [2]int{e[1], blue})
	}

	que := make([][3]int, 0)
	que = append(que, [3]int{0, 0, -1})

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	for len(que) > 0 {
		node, steps, color := que[0][0], que[0][1], que[0][2]

		que = que[1:]

		if res[node] == -1 {
			res[node] = steps
		}

		for i := range gr[node] {
			adj := &gr[node][i]
			if adj[0] != visited && adj[1] != color {
				que = append(que, [3]int{adj[0], steps + 1, adj[1]})
				adj[0] = visited
			}
		}
	}

	return res
}
