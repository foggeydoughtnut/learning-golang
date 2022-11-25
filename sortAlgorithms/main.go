package main

import (
	"fmt"
	"math/rand"
	"sortAlgorithms/bubbleSort"
	"sortAlgorithms/mergeSort"
	"sortAlgorithms/selectionSort"
	"time"
)

// Generates random list with seed of 100
func generateList(numOfValues int) []int {
	rand.Seed(100)
	list := rand.Perm(numOfValues)
	return list
}

func main() {
	numOfValues := 10000

	bubList := generateList(numOfValues)
	startBub := time.Now()
	bubbleSort.BubbleSort(bubList)
	elaspedTimeBub := time.Since(startBub) / 1e6
	fmt.Printf("Time bubble sort took : %dms\n", elaspedTimeBub)

	mergeList := generateList(numOfValues)
	startMerge := time.Now()
	mergeSort.MergeSort(mergeList)
	elaspedTimeMerge := time.Since(startMerge) / 1e6
	fmt.Printf("Time merge sort took : %dms\n", elaspedTimeMerge)

	selectList := generateList(numOfValues)
	startSelect := time.Now()
	selectionSort.SelectionSort(selectList)
	elaspedTimeSelect := time.Since(startSelect) / 1e6
	fmt.Printf("Time selection sort took : %dms\n", elaspedTimeSelect)
}
