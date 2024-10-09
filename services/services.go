package services

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const Length = 25

func GenerateRandomArray() [Length]int {
	var arr [Length]int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < Length; i++ {
		arr[i] = rand.Intn(1000)
	}

	return arr
}

func DeleteFiles() {

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
		return
	}
}

const (
	Reset  = "\033[0m"    
	Red    = "\033[31m" 
	Green  = "\033[32m" 
	Yellow = "\033[33m" 
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func ShowBanner() {

	banner := fmt.Sprintf(`
%s 	##     ##    ###     ######   ########  ####  ##     ##  ######   ########  ######  
  	###   ###   ## ##   ##    ##  ##     ##  ##   ##     ## ##    ##  ##       ##    ## 
  	#### ####  ##   ##  ##        ##     ##  ##   ##     ## ##        ##       ##       
  	## ### ## ##     ## ##   #### ########   ##   ######### ##   #### ######    ######  
  	##     ## ######### ##    ##  ##         ##   ##     ## ##    ##  ##             ## 
  	##     ## ##     ## ##    ##  ##         ##   ##     ## ##    ##  ##       ##    ## 
  	##     ## ##     ##  ######   ##        ####  ##     ##  ######   ########  ######  
  `, Yellow)

	fmt.Println(banner)
}


func ShowMenu() {
	fmt.Println("\n\n")	
	
	fmt.Println("Escolha uma opção:")
	fmt.Println("\n")

	fmt.Println("1 - Quicksort")
	fmt.Println("2 - Mergesort")

	fmt.Println("\n\n")	
	fmt.Println("0 - Sair")
}

