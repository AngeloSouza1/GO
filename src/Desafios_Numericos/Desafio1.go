package main

import (
	"fmt"
)

func main() {
	fmt.Println("Números entre 1 e 100 que são divisíveis por 3:")

	count := 0 
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%3d ", i)
			count++
			if count%10 == 0 { 
				fmt.Println()
			}
		}
	}

	if count%10 != 0 {
		fmt.Println()
	}
}
