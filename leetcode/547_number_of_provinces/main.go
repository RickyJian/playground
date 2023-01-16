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
	arr   []int // arr 長度與城市數量相同，element 代表每個程式的 root
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
	for u.arr[idx] != idx {
		idx = u.arr[idx]
	}
	return idx
}
