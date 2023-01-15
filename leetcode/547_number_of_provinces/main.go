package main

func main() {
	// TODO: implement here
}

func findCircleNum(isConnected [][]int) int {
	uf := new(uf)
	uf.new(isConnected)
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return len(isConnected) - uf.count
}

type uf struct {
	arr   []int
	count int
}

func (u *uf) new(isConnected [][]int) {
	u.arr = make([]int, len(isConnected))
	for i := range isConnected {
		u.arr[i] = i
	}
}

func (u *uf) union(i, j int) {
	findI, findJ := u.find(i), u.find(j)
	if findI != findJ {
		u.arr[findJ] = findI
		u.count++
	}
}

func (u *uf) find(idx int) int {
	for u.arr[idx] != idx {
		idx = u.arr[idx]
	}
	return idx
}
