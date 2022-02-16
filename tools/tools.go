package tools

import (
	"fmt"
	"os"
	"sync"
)

var pairs = map[int]string{3: "Fizz", 5: "Bazz", 7: "Boom"}

func GetOutputForRange(first, last int, wg *sync.WaitGroup, mutexForMap *sync.RWMutex, mutexForSlice *sync.RWMutex, result []string) {
	defer wg.Done()
	for i := first; i <= last; i++ {
		mutexForSlice.Lock()
		result[i-1] = GetOutput(i, mutexForMap)
		mutexForSlice.Unlock()
	}
}

func GetOutput(in int, mutex *sync.RWMutex) string {
	flag := false
	out := ""
	mutex.RLock()
	for k, v := range pairs {
		if in%k == 0 {
			flag = true
			out += v
		}
	}
	mutex.RUnlock()
	if !flag {
		out = fmt.Sprintf("%d", in)
	}
	return out
}

func GetOutputForRangeWithConditions(first, last int, wg *sync.WaitGroup, mutexForSlice *sync.RWMutex, result []string, conditions map[int]string) {
	defer wg.Done()
	for i := first; i <= last; i++ {
		flag := false
		for divider, output := range conditions {
			if i%divider == 0 {
				flag = true
				mutexForSlice.Lock()
				result[i-1] += output
				mutexForSlice.Unlock()
			}
		}
		if !flag {
			mutexForSlice.Lock()
			result[i-1] = fmt.Sprintf("%d", i)
			mutexForSlice.Unlock()
		}
	}
}

func GetOutputWithConditions(in int, conditions map[int]string) string {
	flag := false
	out := ""
	for divider, output := range conditions {
		if in%divider == 0 {
			flag = true
			out += output
		}
	}
	if !flag {
		out = fmt.Sprintf("%d", in)
	}
	return out
}

func PrintOutput(in int) {
	flag := false
	for k, v := range pairs {
		if in%k == 0 {
			flag = true
			fmt.Print(v)
		}
	}
	if !flag {
		fmt.Print(in)
	}
	fmt.Println()
}

func PrintOutputToFile(in int, file *os.File) {
	flag := false
	for k, v := range pairs {
		if in%k == 0 {
			flag = true
			fmt.Fprint(file, v)
		}
	}
	if !flag {
		fmt.Fprint(file, in)
	}
	fmt.Fprint(file, "\n")
}
