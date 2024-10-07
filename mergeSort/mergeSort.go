package mergeSort

import (
	"algoritmos/plot"
)

func MergeSort(arr []int, iteration *int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := MergeSort(arr[:middle], iteration)
	right := MergeSort(arr[middle:], iteration)

	return merge(left, right, iteration)
}

func merge(left []int, right []int, iteration *int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}

		*iteration++
		plot.PlotArray(append(result, append(left[i:], right[j:]...)...), *iteration, "Merge Sort", []int{i, j}) 
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
