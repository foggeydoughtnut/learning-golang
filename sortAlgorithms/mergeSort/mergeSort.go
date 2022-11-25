package mergeSort

func MergeSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	left := MergeSort(values[:len(values)/2])
	right := MergeSort(values[len(values)/2:])
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
