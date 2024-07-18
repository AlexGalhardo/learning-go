package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	secretCodeToFind = 213274
	minNumber        = 100000
	maxNumber        = 999999
)

var (
	totalNumbersToEachWorkerToTest int
	startAt                        []int
	endAt                          []int
)

func initRanges(numWorkers int) {
	startAt = make([]int, numWorkers)
	endAt = make([]int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		if i == 0 {
			startAt[0] = minNumber
			endAt[0] = minNumber + totalNumbersToEachWorkerToTest
		} else {
			startAt[i] = endAt[i-1] + 1
			endAt[i] = startAt[i] + totalNumbersToEachWorkerToTest
			if endAt[i] > maxNumber {
				endAt[i] = maxNumber
			}
		}
	}
}

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func Worker(id, start, end int, wg *sync.WaitGroup, found chan struct{}) {
	defer wg.Done()

	for i := start; i <= end; i++ {
		select {
		case <-found:
			return
		default:
			fmt.Printf("...WORKER ID %d => trying number: %d\n", id, i)
			sleep(25)

			if i == secretCodeToFind {
				fmt.Printf("\n\nTHE SECRET CODE is =========> %d\n\n", i)
				close(found)
				return
			}
		}
	}
}

func main() {
	numWorkers := runtime.NumCPU()
	totalNumbersToEachWorkerToTest = (maxNumber - minNumber) / numWorkers

	initRanges(numWorkers)

	var wg sync.WaitGroup
	found := make(chan struct{})

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go Worker(i, startAt[i], endAt[i], &wg, found)
	}

	wg.Wait()
}
