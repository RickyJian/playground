package main

import (
	"context"
	"errors"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	worker := make(chan struct{}, 50)
	queue := make(chan *result, 100)

	results := make(map[int]interface{}, 500)
	done := make(chan interface{})
	go func() {
		var count int
	LOOP:
		for q := range queue {
			results[q.batch] = q.content
			if q.err != nil || count == 499 {
				cancel()
				close(done)
				break LOOP
			}
			count++
		}
	}()

LOOP:
	for i := 0; i < 500; i++ {
		worker <- struct{}{}
		select {
		case <-ctx.Done():
			break LOOP
		default:
			go func(batch int) {
				select {
				case <-ctx.Done():
					// if error occur do nothing
				default:
					queue <- process(ctx, batch)
				}
				<-worker
			}(i)
		}
	}

	<-done
}

type result struct {
	batch   int
	content interface{}
	err     error
}

func process(ctx context.Context, batch int) *result {
	time.Sleep(time.Second)
	r := &result{}
	select {
	case <-ctx.Done():
		r.err = ctx.Err()
	default:
		r.batch = batch
		r.content = batch
		if batch == 200 {
			r.err = errors.New("error occur")
		}
	}
	return r
}
