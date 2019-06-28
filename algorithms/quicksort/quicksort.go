package main

import "fmt"

func quicksort(array []int, low, high int) {

	if low < high {
		idx := partition(array, low, high)

		go quicksort(array[:idx], 0, len(array[:idx])-1)
		go quicksort(array[idx+1:], 0, len(array[idx+1:])-1)
	}
}

func partition(array []int, low, high int) int {

	pivot := array[high]
	i := low - 1

	for j := low; j < high; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]

	return i + 1

}

func main() {

	input := []int{10, 80, 30, 90, 40, 50, 70}

	quicksort(input, 0, len(input)-1)

	fmt.Println(input)

}
