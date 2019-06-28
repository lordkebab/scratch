package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	mid := len(array) / 2
	return merge(mergeSort(array[:mid]), mergeSort(array[mid:]))
}

func merge(left, right []int) []int {
	var res []int

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	if i < len(left) {
		res = append(res, left[i])
	}

	if j < len(right) {
		res = append(res, right[j])
	}

	return res
}

func main() {
	input := []int{5, 4, 1, 8, 7, 2, 6, 3}

	sorted := mergeSort(input)
	fmt.Println(sorted)
}
