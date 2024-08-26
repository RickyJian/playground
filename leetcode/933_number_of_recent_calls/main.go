package main

func main() {
	// TODO: implement here
}

type RecentCounter struct {
	reqs []int
}

func Constructor() RecentCounter {
	return RecentCounter{reqs: make([]int, 0, 10000)}
}

func (this *RecentCounter) PingV1(t int) int {
	for len(this.reqs) > 0 && t-3000 > this.reqs[0] {
		this.reqs = this.reqs[1:]
	}
	this.reqs = append(this.reqs, t)
	return len(this.reqs)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
