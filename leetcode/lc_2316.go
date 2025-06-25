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
func countPairs1(n int, edges [][]int) int64 {

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

// UnionFind 并查集方法实现
type UnionFind struct {
	num     int
	parents map[int]int
	size    map[int]int
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.num = n
	uf.parents = make(map[int]int)
	uf.size = make(map[int]int)
	for i := 0; i < n; i++ {
		uf.parents[i] = i
		uf.size[i] = 1
	}
	return uf
}
func (uf *UnionFind) find(x int) int {
	if uf.parents[x] == x {
		return x
	}
	//路径压缩
	uf.parents[x] = uf.find(uf.parents[x])
	return uf.parents[x]
}
func (uf *UnionFind) union(x, y int) {

	xp := uf.find(x)
	yp := uf.find(y)
	if xp != yp {
		sizeX := uf.size[xp]
		sizeY := uf.size[yp]
		//小集合合并到大的集合中
		if sizeX >= sizeY {
			uf.parents[yp] = xp
			uf.size[xp] += sizeY
		} else {
			uf.parents[xp] = yp
			uf.size[yp] += sizeX
		}
	}
}

func countPairs(n int, edges [][]int) int64 {

	uf := NewUnionFind(n)
	for _, edge := range edges {
		uf.union(edge[0], edge[1])
	}
	total := int64(0)
	//遍历每个顶点，累加该顶点之外的所有顶点，因为会重复计算，所以最后结果需要除以2
	for i := 0; i < n; i++ {
		parentI := uf.find(i)
		total += int64(n - uf.size[parentI])
	}
	return total / 2
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
