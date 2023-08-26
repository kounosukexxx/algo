package l323

func countComponents(n int, edges [][]int) int {
	graph := initGraph(n, edges)
	return graph.getNumberOfTrees()
}

type graph [][]int

func initGraph(n int, edges [][]int) graph {
	graph := make(graph, n)
	for _, edge := range edges {
		graph.addEdge(edge[0], edge[1])
	}
	return graph
}

func (g graph) addEdge(w, v int) {
	g[w] = append(g[w], v)
	g[v] = append(g[v], w)
}

func (g graph) getNumberOfTrees() int {
	ans := 0
	visited := make([]bool, len(g))
	for i := 0; i < len(g); i++ {
		if visited[i] {
			continue
		}
		ans++
		g.markTreeAsVisited(i, visited)
	}

	return ans
}

func (g graph) markTreeAsVisited(elem int, visited []bool) {
	visited[elem] = true
	for _, v := range g[elem] {
		if visited[v] {
			continue
		}
		g.markTreeAsVisited(v, visited)
	}
}
