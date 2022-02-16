package main

import (
	"fizzBuzz/tools"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	amountOfIterations := 3000000
	var n int
	var fileOutput string
	var mapAllocation string
	fmt.Println("Are you going to use file output?(yes/no)")
	fmt.Scanf("%s\n", &fileOutput)
	fmt.Println("Enter the number of goroutines")
	fmt.Scanf("%d\n", &n)
	if n > 1 {
		fmt.Println("Are you going to allocate map with dividers and output values to each goroutine?(yes/no)")
		fmt.Scanf("%s\n", &mapAllocation)
	}
	fmt.Println()
	start := time.Now()
	if fileOutput == "yes" {
		file, err := os.Create("out.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		switch n {
		case 1:
			for i := 1; i <= amountOfIterations; i++ {
				tools.PrintOutputToFile(i, file)
			}
			break
		default:
			var wg sync.WaitGroup
			var mutexForSlice sync.RWMutex
			stageSize := amountOfIterations / n
			result := make([]string, amountOfIterations)
			if mapAllocation == "yes" {
				for i := 0; i < n-1; i++ {
					wg.Add(1)
					var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
					go tools.GetOutputForRangeWithConditions(stageSize*i+1, stageSize*(i+1), &wg, &mutexForSlice, result, conditions)
				}
				for i := stageSize*(n-1) + 1; i <= amountOfIterations; i++ {
					var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
					mutexForSlice.Lock()
					result[i-1] = tools.GetOutputWithConditions(i, conditions)
					mutexForSlice.Unlock()
				}
			} else {
				var mutexForMap sync.RWMutex
				for i := 0; i < n-1; i++ {
					wg.Add(1)
					go tools.GetOutputForRange(stageSize*i+1, stageSize*(i+1), &wg, &mutexForMap, &mutexForSlice, result)
				}
				for i := stageSize*(n-1) + 1; i <= amountOfIterations; i++ {
					mutexForSlice.Lock()
					result[i-1] = tools.GetOutput(i, &mutexForMap)
					mutexForSlice.Unlock()
				}
			}
			wg.Wait()
			for _, v := range result {
				fmt.Fprintln(file, v)
			}
		}
	} else {
		switch n {
		case 1:
			for i := 1; i <= amountOfIterations; i++ {
				tools.PrintOutput(i)
			}
			break
		default:
			var wg sync.WaitGroup
			var mutexForSlice sync.RWMutex
			stageSize := amountOfIterations / n
			result := make([]string, amountOfIterations)
			if mapAllocation == "yes" {
				for i := 0; i < n-1; i++ {
					wg.Add(1)
					var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
					go tools.GetOutputForRangeWithConditions(stageSize*i+1, stageSize*(i+1), &wg, &mutexForSlice, result, conditions)
				}
				for i := stageSize*(n-1) + 1; i <= amountOfIterations; i++ {
					var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
					mutexForSlice.Lock()
					result[i-1] = tools.GetOutputWithConditions(i, conditions)
					mutexForSlice.Unlock()
				}
			} else {
				var mutexForMap sync.RWMutex
				for i := 0; i < n-1; i++ {
					wg.Add(1)
					go tools.GetOutputForRange(stageSize*i+1, stageSize*(i+1), &wg, &mutexForMap, &mutexForSlice, result)
				}
				for i := stageSize*(n-1) + 1; i <= amountOfIterations; i++ {
					mutexForSlice.Lock()
					result[i-1] = tools.GetOutput(i, &mutexForMap)
					mutexForSlice.Unlock()
				}
			}
			wg.Wait()
			for _, v := range result {
				fmt.Println(v)
			}
		}
	}
	end := time.Now()
	fmt.Printf("Computation has taken %d nanoseconds\n", end.Sub(start).Nanoseconds())
}
