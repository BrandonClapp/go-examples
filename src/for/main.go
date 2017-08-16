package main

import (
	"fmt"
)

func main() {
	doForConditional()
	doForInfinite()
	doForExpression()
}

func doForConditional() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func doForInfinite() {
	i := 1
	for {
		if i <= 10 {
			fmt.Println(i)
			i++
			continue
		}
		break
	}
}

func doForExpression() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
