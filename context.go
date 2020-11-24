package main

import (
	"fmt"
	"sync"
	"time"
	"context"
)

const duration = 5
const timeoutDuration = time.Second * duration

func main() {

	// create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("Connecting to server...")
		if err := connect(ctx); err != nil {
			fmt.Println(err)

			// if connection times out, cancel everything
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		
		fmt.Println("Loading assets...")
		if err := load(ctx); err != nil {
			fmt.Println(err)
		}

	}()

	wg.Wait()
}

func connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)

	select {
	case <-ctx.Done():
		cancel()
		return ctx.Err()
	}

	return nil
}

func load(ctx context.Context) error {
	loads := []string{
		"Account status",
		"Account balance",
		"Code updates",
		"Orderbook",
	}

	for _, v := range loads {
		// simulate loading time
		select {
		case <-time.After(time.Second * 3):
			fmt.Println(v)
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}
