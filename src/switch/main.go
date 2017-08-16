package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		sayNumber(i)
	}
}

func sayNumber(i int) {
	switch i {
	case 0:
		{
			// multiple lines with curly braces
			fmt.Println("Zero")
		}
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}
}
