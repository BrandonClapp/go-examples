package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// cd src/goroutine1 && time go run main.go

	names := []string{ "Brandon", "Barabosa", "Jeffery" }

	// concurrent sequential: time GOMAXPROCS=1 go run main.go
	// concurrent parallel: time go run main.go

	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names {
		go doCalculation(name, &wg)
	}
	wg.Wait()
}

func doCalculation(n string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println(n)
	wg.Done()
}