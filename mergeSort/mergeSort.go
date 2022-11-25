package main

import "fmt"

func mergeSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	left := mergeSort(values[:len(values)/2])
	right := mergeSort(values[len(values)/2:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	sorted := []int{}
	left_index := 0
	right_index := 0
	for left_index < len(left) && right_index < len(right) {
		if left[left_index] < right[right_index] {
			sorted = append(sorted, left[left_index])
			left_index++
		} else {
			sorted = append(sorted, right[right_index])
			right_index++
		}
	}
	for ; left_index < len(left); left_index++ {
		sorted = append(sorted, left[left_index])
	}
	for ; right_index < len(right); right_index++ {
		sorted = append(sorted, right[right_index])
	}
	return sorted
}

func main() {
	list := []int{1, 6, 2, 4, 98, 0, 1, 3, 5, 7, 6}
	fmt.Printf("Values before sort : %d\n", list)
	values := mergeSort(list)
	fmt.Printf("Values after sort : %d\n", values)
}
