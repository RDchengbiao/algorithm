package main

import (
	"sort"
)

// 406. 根据身高重建队列

type People struct {
	Height int
	K      int
}

type Peoples []People

func (p Peoples) Len() int {
	return len(p)
}
func (p Peoples) Less(i, j int) bool {
	if p[i].Height == p[j].Height {
		return p[i].K > p[j].K
	}
	return p[i].Height < p[j].Height
}
func (p Peoples) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 先按照身高从小到大排序，身高相同，则按照k值从大到小排序，然后从结果数组中找空位置，k值是多少，空位置就需要预留出来多少个。
// 身高排序后，前面的矮子影响不到当前的人顺序，当前人只需要保证前面有k个空位即可保证自己的k值匹配，对于身高相同的人来说，先安排好
// k值大的，k值小的放到他前面也不会影响高k值的k值的正确性。
func reconstructQueue(people [][]int) [][]int {

	peoples := make(Peoples, len(people))
	for i := 0; i < len(people); i++ {
		peoples[i] = People{
			Height: people[i][0],
			K:      people[i][1],
		}
	}
	n := len(people)
	res := make([][]int, len(people))
	sort.Sort(peoples)
	cnt := 0
	for _, peop := range peoples {
		cnt = peop.K
		j := 0
		for ; j < n; j++ {
			if len(res[j]) == 0 {
				break
			}
		}
		for ; j < n; j++ {
			if cnt == 0 && len(res[j]) == 0 {
				break
			}
			if len(res[j]) == 0 {
				cnt--
			}
		}
		res[j] = []int{peop.Height, peop.K}
	}
	return res
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	reconstructQueue(people)
}
