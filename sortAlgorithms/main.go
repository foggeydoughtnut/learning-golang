package main

import (
	"fmt"
	"math/rand"
	"sortAlgorithms/bubbleSort"
	"sortAlgorithms/mergeSort"
	"sortAlgorithms/selectionSort"
	"time"
)

func main() {
	rand.Seed(100)
	numOfValues := 10000

	bubList := rand.Perm(numOfValues)
	startBub := time.Now()
	bubbleSort.BubbleSort(bubList)
	elaspedTimeBub := time.Since(startBub) / 1e6
	fmt.Printf("Time bubble sort took : %dms\n", elaspedTimeBub)

	mergeList := rand.Perm(numOfValues)
	startMerge := time.Now()
	mergeSort.MergeSort(mergeList)
	elaspedTimeMerge := time.Since(startMerge) / 1e6
	fmt.Printf("Time merge sort took : %dms\n", elaspedTimeMerge)

	selectList := rand.Perm(numOfValues)
	startSelect := time.Now()
	selectionSort.SelectionSort(selectList)
	elaspedTimeSelect := time.Since(startSelect) / 1e6
	fmt.Printf("Time selection sort took : %dms\n", elaspedTimeSelect)
}
