package main

import (
	"fmt"
	"skn-go-practice/src2/pkgs/utils"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = utils.GetWaitGroup()
	var mutex *sync.Mutex = utils.GetMutex()

	var score []int = []int{
		0,
	}

	wg.Add(3)

	go func() {
		fmt.Println("First One")

		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()

		wg.Done()
	}()

	go func() {
		fmt.Println("Second One")

		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()

		wg.Done()
	}()

	go func() {
		fmt.Println("Third One")

		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Score:", score)
}
