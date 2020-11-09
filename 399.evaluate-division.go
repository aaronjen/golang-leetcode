package main

/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
func calcEquationSolve(graph map[string]map[string]float64, start string, end string, visited map[string]bool) float64 {
	if _, ok := graph[start]; !ok {
		return -1
	}

	if v, ok := graph[start][end]; ok {
		return v
	}

	visited[start] = true
	for neighbor, weight := range graph[start] {
		if !visited[neighbor] {
			visited[neighbor] = true
			v := calcEquationSolve(graph, neighbor, end, visited)

			if v != -1 {
				graph[start][end] = weight * v
				return weight * v
			}
		}
	}

	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]map[string]float64{}

	n := len(equations)

	for i := 0; i < n; i++ {
		start, end := equations[i][0], equations[i][1]

		if _, ok := graph[start]; !ok {
			graph[start] = map[string]float64{}
		}
		if _, ok := graph[end]; !ok {
			graph[end] = map[string]float64{}
		}

		graph[start][end] = values[i]
		graph[end][start] = 1 / values[i]
	}

	res := make([]float64, len(queries))

	for i := 0; i < len(queries); i++ {
		start, end := queries[i][0], queries[i][1]
		visited := map[string]bool{}
		res[i] = calcEquationSolve(graph, start, end, visited)
	}

	return res

}

// @lc code=end
