package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate some random numbers
	generator := func(done <-chan interface{}) chan int {
		// channel to send random numbers on
		c := make(chan int)
		go func() {
			defer close(c)
			for {
				select {
				case <-done:
					return
				case c <- rand.Int():
				}
			}
		}()
		return c
	}

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	done := make(chan interface{})

	// read some random numbers
	g := generator(done)
	for i:=0; i<3; i++ {
		fmt.Println(<-g)
	}

	fmt.Println("-------")

	// read some more random numbers
	for j:=0; j<5; j++ {
		fmt.Println(<-g)
	}

	close(g)
}
