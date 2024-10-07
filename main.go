package main

import (
	"algoritmos/plot" 
	// "algoritmos/quickSort"
	"algoritmos/mergeSort" 
	"fmt"
	"os"
	"path/filepath"
	"math/rand"
	"time"
)

func deleteFiles() {
	dir := "graphics" // Altere para o caminho do seu diretório

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() { // Verifica se não é um diretório
			return os.Remove(path) // Remove o arquivo
		}
		return nil
	})

	if err != nil {
		fmt.Println("Erro ao apagar os arquivos:", err)
		return
	}

	fmt.Println("Todos os arquivos foram apagados com sucesso.")
}

func generateRandomArray() [25]int {
	var arr [25]int
	rand.Seed(time.Now().UnixNano()) 

	for i := 0; i < 25; i++ {
		arr[i] = rand.Intn(1000) 
	}

	return arr
}


func main() {
	// name := "QuickSort"
	name := "MergeSort"

	arr := generateRandomArray()
	iteration := 0

	deleteFiles()
	plot.PlotArray(arr[:], iteration, name, []int{})

	// quickSort.QuickSort(arr[:], 0, len(arr)-1, &iteration)
	mergeSort.MergeSort(arr[:], &iteration)

	fmt.Println("Ordenação final:", arr)
}