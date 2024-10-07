package quickSort

import (
	"algoritmos/plot"
)

func Partition(arr []int, low, high int, iteration *int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			*iteration++
			plot.PlotArray(arr, *iteration)
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	*iteration++
	plot.PlotArray(arr, *iteration)
	return i + 1
}

func QuickSort(arr []int, low, high int, iteration *int) {
	if low < high {
		pi := Partition(arr, low, high, iteration)
		QuickSort(arr, low, pi-1, iteration)
		QuickSort(arr, pi+1, high, iteration)
	}
}