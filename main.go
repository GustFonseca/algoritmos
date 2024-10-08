package main

import (
	"algoritmos/plot" 
	"algoritmos/quickSort"
	// "algoritmos/mergeSort" 
	"fmt"
	"os"
	"path/filepath"
	"math/rand"
	"time"
)

func deleteFiles() {
	dir := "graphics" 

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() { 
			return os.Remove(path) 
		}
		return nil
	})

	if err != nil {
		fmt.Println("Erro ao apagar os arquivos:", err)
		return
	}

	fmt.Println("Todos os arquivos foram apagados com sucesso.")
}

func generateRandomArray() [100]int {
	var arr [100]int
	rand.Seed(time.Now().UnixNano()) 

	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(1000) 
	}

	return arr
}


func main() {
	name := "QuickSort"
	// name := "MergeSort"

	arr := generateRandomArray()
	iteration := 0

	deleteFiles()
	plot.PlotArray(arr[:], iteration, name, []int{})

	quickSort.QuickSort(arr[:], 0, len(arr)-1, &iteration)
	// mergeSort.MergeSort(arr[:], &iteration)

	fmt.Println("Ordenação final:", arr)
}