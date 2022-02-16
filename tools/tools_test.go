package tools

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func BenchmarkPrintOutputToFile_byOneGoroutine(b *testing.B) {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	amountOfIterations := 1000
	for i := 0; i < b.N; i++ {
		for j := 1; j < amountOfIterations; j++ {
			PrintOutputToFile(j, file)
		}
	}
}

func BenchmarkPrintOutputToFile_byTwoGoroutines(b *testing.B) {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	amountOfIterations := 1000
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var mutexForMap sync.RWMutex
		var mutexForSlice sync.RWMutex
		stageSize := amountOfIterations / 2
		result := make([]string, amountOfIterations)
		wg.Add(1)
		go GetOutputForRange(1, stageSize, &wg, &mutexForMap, &mutexForSlice, result)
		for j := stageSize + 1; j <= amountOfIterations; j++ {
			mutexForSlice.Lock()
			result[j-1] = GetOutput(j, &mutexForMap)
			mutexForSlice.Unlock()
		}
		wg.Wait()
		for _, v := range result {
			fmt.Fprintln(file, v)
		}
	}
}

func BenchmarkPrintOutputToFile_byThreeGoroutines(b *testing.B) {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	amountOfIterations := 1000
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var mutexForMap sync.RWMutex
		var mutexForSlice sync.RWMutex
		stageSize := amountOfIterations / 3
		result := make([]string, amountOfIterations)
		for j := 0; j < 2; j++ {
			wg.Add(1)
			go GetOutputForRange(stageSize*j+1, stageSize*(j+1), &wg, &mutexForMap, &mutexForSlice, result)
		}
		for j := stageSize*2 + 1; j <= amountOfIterations; j++ {
			mutexForSlice.Lock()
			result[j-1] = GetOutput(j, &mutexForMap)
			mutexForSlice.Unlock()
		}
		wg.Wait()
		for _, v := range result {
			fmt.Fprintln(file, v)
		}
	}
}

func BenchmarkPrintOutputToFile_byFourGoroutines(b *testing.B) {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	amountOfIterations := 1000
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var mutexForMap sync.RWMutex
		var mutexForSlice sync.RWMutex
		stageSize := amountOfIterations / 4
		result := make([]string, amountOfIterations)
		for j := 0; j < 3; j++ {
			wg.Add(1)
			go GetOutputForRange(stageSize*j+1, stageSize*(j+1), &wg, &mutexForMap, &mutexForSlice, result)
		}
		for j := stageSize*3 + 1; j <= amountOfIterations; j++ {
			mutexForSlice.Lock()
			result[j-1] = GetOutput(j, &mutexForMap)
			mutexForSlice.Unlock()
		}
		wg.Wait()
		for _, v := range result {
			fmt.Fprintln(file, v)
		}
	}
}

func BenchmarkPrintOutputToFileWithConditions_byTwoGoroutines(b *testing.B) {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	amountOfIterations := 1000
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var mutexForSlice sync.RWMutex
		stageSize := amountOfIterations / 2
		result := make([]string, amountOfIterations)
		wg.Add(1)
		var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
		go GetOutputForRangeWithConditions(1, stageSize, &wg, &mutexForSlice, result, conditions)
		for j := stageSize + 1; j <= amountOfIterations; j++ {
			var conditions2 = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
			mutexForSlice.Lock()
			result[j-1] = GetOutputWithConditions(j, conditions2)
			mutexForSlice.Unlock()
		}
		wg.Wait()
		for _, v := range result {
			fmt.Fprintln(file, v)
		}
	}
}

func BenchmarkPrintOutputToFileWithConditions_byThreeGoroutines(b *testing.B) {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	amountOfIterations := 1000
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var mutexForSlice sync.RWMutex
		stageSize := amountOfIterations / 3
		result := make([]string, amountOfIterations)
		for j := 0; j < 2; j++ {
			wg.Add(1)
			var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
			go GetOutputForRangeWithConditions(stageSize*j+1, stageSize*(j+1), &wg, &mutexForSlice, result, conditions)
		}
		for j := stageSize*2 + 1; j <= amountOfIterations; j++ {
			var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
			mutexForSlice.Lock()
			result[j-1] = GetOutputWithConditions(j, conditions)
			mutexForSlice.Unlock()
		}
		wg.Wait()
		for _, v := range result {
			fmt.Fprintln(file, v)
		}
	}
}

func BenchmarkPrintOutputToFileWithConditions_byFourGoroutines(b *testing.B) {
	file, err := os.Create("out.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	amountOfIterations := 1000
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var mutexForSlice sync.RWMutex
		stageSize := amountOfIterations / 4
		result := make([]string, amountOfIterations)
		for j := 0; j < 3; j++ {
			wg.Add(1)
			var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
			go GetOutputForRangeWithConditions(stageSize*j+1, stageSize*(j+1), &wg, &mutexForSlice, result, conditions)
		}
		for j := stageSize*3 + 1; j <= amountOfIterations; j++ {
			var conditions = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}
			mutexForSlice.Lock()
			result[j-1] = GetOutputWithConditions(j, conditions)
			mutexForSlice.Unlock()
		}
		wg.Wait()
		for _, v := range result {
			fmt.Fprintln(file, v)
		}
	}
}

func TestGetOutputForRange(t *testing.T) {
	samples := []struct {
		test               string
		first              int
		last               int
		numberOfGoroutines int
		result             []string
		expected           []string
	}{
		{
			test:               "[1:7]",
			first:              1,
			last:               7,
			numberOfGoroutines: 2,
			result:             make([]string, 7),
			expected:           []string{"1", "2", "Fizz", "4", "Bazz", "Fizz", "Boom"},
		},
		{
			test:               "[9:13]",
			first:              9,
			last:               13,
			numberOfGoroutines: 3,
			result:             make([]string, 13),
			expected:           []string{"Fizz", "Bazz", "11", "Fizz", "13"},
		},
	}
	for _, sample := range samples {
		t.Run(sample.test, func(t *testing.T) {
			var wg sync.WaitGroup
			var mutexForMap sync.RWMutex
			var mutexForSlice sync.RWMutex
			stageSize := sample.first / sample.numberOfGoroutines
			for i := 0; i < sample.numberOfGoroutines-1; i++ {
				wg.Add(1)
				go GetOutputForRange(stageSize*i+1, stageSize*(i+1), &wg, &mutexForMap, &mutexForSlice, sample.result)
			}
			wg.Add(1)
			go GetOutputForRange(stageSize*(sample.numberOfGoroutines-1)+1, sample.last, &wg, &mutexForMap, &mutexForSlice, sample.result)
			wg.Wait()
			for i, expectedOutput := range sample.expected {
				if sample.result[i+sample.first-1] != expectedOutput {
					t.Fatalf("\t%s: %s", failed, fmt.Errorf("Index %d: expected %s, got %s ", i, expectedOutput, sample.result[i+sample.first-1]))
				}
			}
			t.Logf("\t\t%s", succeed)
		})
	}
}

func TestGetOutput(t *testing.T) {
	samples := []struct {
		test     string
		in       int
		expected string
	}{
		{
			test:     "Input: 1",
			in:       1,
			expected: "1",
		},
		{
			test:     "Input: 15",
			in:       15,
			expected: "FizzBazz",
		},
		{
			test:     "Input: 21",
			in:       21,
			expected: "FizzBoom",
		},
	}
	for _, sample := range samples {
		t.Run(sample.test, func(t *testing.T) {
			var mutex sync.RWMutex
			out := GetOutput(sample.in, &mutex)
			if out != sample.expected {
				t.Fatalf("\t%s: %s", failed, fmt.Errorf("Expected %s, got %s ", sample.expected, out))
			}
			t.Logf("\t\t%s", succeed)
		})
	}
}
