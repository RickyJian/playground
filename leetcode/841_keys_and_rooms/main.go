package main

func main() {
	// TODO: implement here
}

func canVisitAllRoomsV1(rooms [][]int) bool {
	visited := map[int]struct{}{0: {}}
	keys := rooms[0]
	for len(keys) > 0 {
		first := keys[0]
		if _, ok := visited[first]; !ok {
			for _, k := range rooms[first] {
				keys = append(keys, k)
			}
			visited[first] = struct{}{}
		}
		keys = keys[1:]
	}
	return len(visited) == len(rooms)
}
