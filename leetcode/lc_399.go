package main

import "fmt"

// 399. 除法求值

type Edge struct {
	start, end string
	weight     float64
}

// todo  解法还有bfs 、floyd 或者 带权并查集
// dfs
func find(graph map[string][]Edge, vis map[string]bool, s, e string, weight float64) float64 {
	if s == e {
		return weight
	}
	vis[s] = true
	for _, edge := range graph[s] {
		if _, ok := graph[edge.end]; ok {
			if vis[edge.end] {
				continue
			}
			tmpRes := find(graph, vis, edge.end, e, weight*edge.weight)
			if tmpRes != -1 {
				vis[s] = false
				return tmpRes
			}
		}
	}
	vis[s] = false
	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	graph := make(map[string][]Edge)
	mp := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		if _, ex := graph[equations[i][0]]; !ex {
			graph[equations[i][0]] = make([]Edge, 0)
		}
		if _, ex := graph[equations[i][1]]; !ex {
			graph[equations[i][1]] = make([]Edge, 0)
		}
		if _, ex := mp[equations[i][0]]; !ex {
			mp[equations[i][0]] = make(map[string]float64)
		}
		if _, ex := mp[equations[i][1]]; !ex {
			mp[equations[i][1]] = make(map[string]float64)
		}
		graph[equations[i][0]] = append(graph[equations[i][0]], Edge{
			start:  equations[i][0],
			end:    equations[i][1],
			weight: values[i],
		})
		graph[equations[i][1]] = append(graph[equations[i][1]], Edge{
			start:  equations[i][1],
			end:    equations[i][0],
			weight: 1.0 / values[i],
		})
		mp[equations[i][0]][equations[i][1]] = values[i]
		mp[equations[i][1]][equations[i][0]] = 1.0 / values[i]
	}
	vis := make(map[string]bool)
	res := make([]float64, 0)
	for i := 0; i < len(queries); i++ {
		s, e := queries[i][0], queries[i][1]
		if _, ex := graph[s]; !ex {
			res = append(res, -1)
			continue
		}
		if _, ex := graph[e]; !ex {
			res = append(res, -1)
			continue
		}
		if _, ex := mp[s]; !ex {
			mp[s] = make(map[string]float64)
		} else {
			if v, ex := mp[s][e]; ex {
				res = append(res, v)
				continue
			}
		}

		ret := find(graph, vis, s, e, 1)
		if ret == -1 {
			res = append(res, -1)
		} else {
			if _, ex := mp[s][e]; !ex {
				mp[s][e] = ret
				mp[e][s] = 1.0 / ret
			}
			graph[s] = append(graph[s], Edge{
				start:  s,
				end:    e,
				weight: ret,
			})

			graph[e] = append(graph[e], Edge{
				start:  e,
				end:    s,
				weight: 1.0 / ret,
			})
			res = append(res, ret)
		}

	}
	return res
}
func main() {
	/*
		equations := [][]string{{"a", "b"}, {"b", "c"}}
		values := []float64{2.0, 3.0}
		queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	*/
	equations := [][]string{{"a", "b"}, {"a", "c"}, {"d", "e"}, {"d", "f"}, {"a", "d"}, {"aa", "bb"}, {"aa", "cc"}, {"dd", "ee"}, {"dd", "ff"}, {"aa", "dd"}, {"a", "aa"}}
	//[["a","b"],["a","c"],["d","e"],["d","f"],["a","d"],["aa","bb"],["aa","cc"],["dd","ee"],["dd","ff"],["aa","dd"],["a","aa"]]

	values := []float64{2.0, 3.0, 4.0, 5.0, 7.0, 5.0, 8.0, 9.0, 3.0, 2.0, 2.0}
	queries := [][]string{{"ff", "a"}}

	res := calcEquation(equations, values, queries)
	fmt.Println(res)

}
