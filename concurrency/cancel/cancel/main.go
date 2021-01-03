package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

const localeUS = "EN/US"

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			cancel()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
			cancel()
		}
	}()
	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world\n", greeting)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == localeUS:
		return "hello", nil
	}
	return "", errors.New("unsupported locale")
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world\n", farewell)
	return nil
}

func genFarewell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == localeUS:
		return "goodbye", nil
	}
	return "", errors.New("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(time.Minute):
	}
	return "EN/US", nil
}
