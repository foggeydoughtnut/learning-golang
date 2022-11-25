package selectionSort

func SelectionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		index := i
		for j := i; j < len(list); j++ {
			if list[j] < list[index] {
				index = j
			}
		}
		list[i], list[index] = list[index], list[i]
	}
	return list
}
