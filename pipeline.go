package main

import (
	"fmt"
)

func main() {

	generator := func(done <-chan interface{}, ints ...int) <-chan int{
		// make a channel for ints to be written to
		c := make(chan int)

		go func() {
			defer fmt.Println("generator done")
			defer close(c)
			// range over the ints and write to the channel
			for _, i := range ints {
				select {
					case c <- i:
					case <-done:
						return
				}
			}
		}()
		return c
	}

	multiply := func(done <-chan interface{}, intStream <-chan int) <-chan int {
		// make a channel for multiplied ints to be written to
		m := make(chan int)

		go func() {
			defer fmt.Println("multiply done")
			defer close(m)
			// range over the ints and multiply
			for i := range intStream {
				select {
				case m <- i*2:
				case <-done:
						return
				}
			}
		}()
		return m
	}

	done := make(chan interface{})
	mult := multiply(done, generator(done, 1, 2, 3, 4, 5))

	for g := range mult {
		fmt.Println(g)
	}
	close(done)
}
