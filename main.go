package main

import (
	"algoritmos/services"
	"algoritmos/plot" 
	"algoritmos/quickSort"
	"algoritmos/mergeSort" 
	"fmt"
	"os"
)

func main() {
	arr := services.GenerateRandomArray()
	services.DeleteFiles()

	services.ShowBanner()
	services.ShowMenu()
	
	iteration := 0

	var choice int
	fmt.Scan(&choice)
	
	switch choice {
	case 1:
		fmt.Println("Quicksort!")
		fmt.Println("\n")
		
		
		plot.PlotArray(arr[:], iteration, "Quicksort", []int{})
		quickSort.QuickSort(arr[:], 0, len(arr)-1, &iteration)

		services.RunShellScript("createGif.sh")

	case 2:
		fmt.Println("Mergesort!")
		fmt.Println("\n")
		
		
		plot.PlotArray(arr[:], iteration, "Mergesort", []int{})
		mergeSort.MergeSort(arr[:], &iteration)

		services.RunShellScript("createGif.sh")

	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida. Tente novamente.")
	}
	
	fmt.Println("Ordenação final:", arr)
}