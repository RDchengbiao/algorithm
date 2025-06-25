package main

//2316. 统计无向图中无法互相到达点对数

func dfs(graph map[int][]int, start int, vis map[int]bool, cnt *int) {
	if vis[start] {
		return
	}
	vis[start] = true
	*cnt++
	for _, end := range graph[start] {
		if !vis[end] {
			dfs(graph, end, vis, cnt)
		}
	}
}
func countPairs(n int, edges [][]int) int64 {

	graph := make(map[int][]int)

	//建图
	for i := 0; i < len(edges); i++ {
		start := edges[i][0]
		end := edges[i][1]
		if _, ok := graph[start]; !ok {
			graph[start] = make([]int, 0)
		}
		graph[start] = append(graph[start], end)
		if _, ok := graph[end]; !ok {
			graph[end] = make([]int, 0)
		}
		graph[end] = append(graph[end], start)
	}
	vis := make(map[int]bool)
	total, sum := int64(0), int64(0)
	//dfs统计每个连通图的节点数量
	for i := 0; i < n; i++ {
		cnt := 0
		if _, ex := vis[i]; ex {
			continue
		}
		dfs(graph, i, vis, &cnt)

		//通过累加每个连通图的顶点数量，降低计算无法互相到达点对数量的复杂度
		sum += total * int64(cnt)
		total += int64(cnt)
	}

	//累计所有无法互相到达点对数
	return sum
}

func main() {
	edges := make([][]int, 0)
	//[0,2],[0,5],[2,4],[1,6],[5,4]
	edges = append(edges, []int{0, 2})
	edges = append(edges, []int{0, 5})
	edges = append(edges, []int{2, 4})
	edges = append(edges, []int{1, 6})
	edges = append(edges, []int{5, 4})

	countPairs(7, edges)
}
