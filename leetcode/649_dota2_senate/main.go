package main

func main() {
	// TODO: implement here
}

func predictPartyVictoryV1(senate string) string {
	var rq, dq []int
	for i, word := range senate {
		if word == 'R' {
			rq = append(rq, i)
		} else {
			dq = append(dq, i)
		}
	}
	count := len(senate)
	for i := 0; ; i++ {
		if len(rq) == 0 {
			return "Dire"
		} else if len(dq) == 0 {
			return "Radiant"
		}

		r := rq[0]
		rq = rq[1:]
		d := dq[0]
		dq = dq[1:]
		if r < d {
			rq = append(rq, count)
		} else {
			dq = append(dq, count)
		}
		count++
	}
}
