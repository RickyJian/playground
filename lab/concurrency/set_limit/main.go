package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, gctx := errgroup.WithContext(context.Background())

	g.SetLimit(3)

	g.Go(func() error {
		fmt.Println("basic goroutine 1")
		time.Sleep(time.Second * 10)
		return nil
	})

	g.Go(func() error {
		fmt.Println("basic goroutine 2")
		time.Sleep(time.Second * 10)
		return nil
	})

	for i := 0; i < 3; i++ {
		g.Go(func() error {
			fmt.Printf("loop count: %d\n", i)
			nestedFunc(gctx, i)
			return nil
		})
	}

	g.Wait()
	fmt.Println("done")
}

func nestedFunc(ctx context.Context, ii int) {
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		fmt.Println("nested goroutine 1")
		time.Sleep(time.Second * 1)
		return nil
	})

	g.Go(func() error {
		fmt.Println("nested goroutine 2")
		time.Sleep(time.Second * 2)
		return nil
	})

	fmt.Println("wait nested goroutine")
	g.Wait()

	fmt.Println("begin g2")
	g2, _ := errgroup.WithContext(ctx)

	g2.SetLimit(1)

	for i := 100; i < 110; i++ {
		g2.Go(func() error {
			fmt.Printf("inner loop count: %d, %d\n", ii, i)
			time.Sleep(time.Second * 1)
			return nil
		})
	}

	g2.Wait()
}
