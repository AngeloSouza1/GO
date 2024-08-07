package main

import (
	"fmt"
)

const (
	Reset      = "\033[0m"
	Yellow     = "\033[33m"
	LightBlue  = "\033[36m"
)

func main() {
	count := 0 

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(Yellow + "Pin " + LightBlue + "Pan" + Reset)
		} else if i%3 == 0 {
			fmt.Print(Yellow + "Pin" + Reset)
		} else if i%5 == 0 {
			fmt.Print(LightBlue + "Pan" + Reset)
		} else {
			fmt.Print(i)
		}

		count++
		if count%10 == 0 {
			fmt.Println() 
		} else {
			fmt.Print(" , ") 
		}
	}

	
	if count%10 != 0 {
		fmt.Println()
	}
}
