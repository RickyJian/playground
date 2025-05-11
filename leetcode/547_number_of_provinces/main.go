package main

func main() {
	// TODO: implement here
}

func findCircleNum(isConnected [][]int) int {
	u := newUF(isConnected)
	for i := 0; i < len(isConnected); i++ {
		// i+1：當 i==j 代表為相同的城市也就代表是相同的省份因此不需要特別處理
		for j := i + 1; j < len(isConnected); j++ {
			// 若兩城市相連需要將兩城市合併
			if isConnected[i][j] == 1 {
				u.union(i, j)
			}
		}
	}
	return u.count
}

type uf struct {
	arr   []int // arr 長度與城市數量相同，element 代表每個城市的 root
	rank  []int // rank 長度與城市數量相同，預設 rank 為 1
	count int   // 總共省份數量
}

// 1. 初始所有的城市 root，將所有城市的 root 設定為 index
// 2. 將 count 設定為所有城市的數量
func newUF(isConnected [][]int) *uf {
	u := &uf{
		arr:   make([]int, len(isConnected)),
		count: len(isConnected),
	}
	for i := range isConnected {
		u.arr[i] = i
	}
	return u
}

func (u *uf) union(i, j int) {
	// 找出 i 城市與 j 城市的 root，
	findI, findJ := u.find(i), u.find(j)
	// 若 root 不同需要將 root 更新成相同
	if findI != findJ {
		// 將 j 的 root 更新成 i 的 root
		u.arr[findJ] = findI
		// 若合併成功代表省份相同，須扣除 1
		u.count--
	}
}

func (u *uf) find(idx int) int {
	// 找到第 idx 城市的 root
	if u.arr[idx] != idx {
		idx = u.arr[idx]
	}
	return idx
}

func findCircleNumV2(isConnected [][]int) int {
	u := newUFV2(isConnected)
	for i := 0; i < len(isConnected); i++ {
		// i+1：當 i==j 代表為相同的城市也就代表是相同的省份因此不需要特別處理
		for j := i + 1; j < len(isConnected); j++ {
			// 若兩城市相連需要將兩城市合併
			if isConnected[i][j] == 1 {
				u.unionRank(i, j)
			}
		}
	}
	return u.count
}

// 1. 初始所有的城市 root，將所有城市的 root 設定為 index
// 2. 將 count 設定為所有城市的數量
func newUFV2(isConnected [][]int) *uf {
	u := &uf{
		arr:   make([]int, len(isConnected)),
		rank:  make([]int, len(isConnected)),
		count: len(isConnected),
	}
	for i := range isConnected {
		u.arr[i] = i
		// 初始 rank 為 1
		u.rank[i] = 1
	}
	return u
}

// merge rank: 將較小的 rank 合併到較大的 rank
func (u *uf) unionRank(i, j int) {
	findI, findJ := u.findPath(i), u.findPath(j)
	if findI == findJ {
		return
	}

	if u.rank[findI] < u.rank[findJ] {
		u.arr[findI] = findJ
	} else if u.rank[findI] > u.rank[findJ] {
		u.arr[findJ] = findI
	} else {
		// 將 arr[j] 的 root 設定為 arr[i]
		u.arr[findJ] = findI
		// 當 rank 相同時, 代表 j 會串接在 i 的下方，因此 i 的 rank 需要 +1
		u.rank[findI]++
	}
	u.count--
}

// path compression: 將走過的路徑直接指向 root
func (u *uf) findPath(idx int) int {
	if u.arr[idx] != idx {
		u.arr[idx] = u.findPath(u.arr[idx])
	}
	return u.arr[idx]
}
