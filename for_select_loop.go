package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan interface{})

	for _, i := range []int{1, 2, 3, 4, 5} {

		if i == 5 {
			done <- 6
		}
		select {
		case <-done:
			fmt.Println("Done")
			return
		case c <- i:
		}
	}

}
