package main

import (
	"fmt"
	"skn-go-practice/src3/pkgs/utils"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = utils.GetWaitGroup()

	var myChannel chan int = make(chan int)

	wg.Add(2)

	//* By default all goroutines can communicate multi-directional i.e. they can both be receiver & sender
	//* But it can be restricted with a specific way.

	//* receiver
	//* Receiver Channel can't close the channel
	go func(channel <-chan int) {
		//close(myChannel)
		value, isChannelOpen := <-myChannel

		fmt.Printf("Channel Open?: %v \n", isChannelOpen)
		fmt.Printf("My Channel Data: %d \n", value)

		value2, isChannelOpen2 := <-myChannel
		fmt.Printf("Channel Open?: %v \n", isChannelOpen2)
		fmt.Printf("My Channel Data2: %d \n", value2)

		value3, isChannelOpen3 := <-myChannel
		fmt.Printf("Channel Open?: %v \n", isChannelOpen3)
		fmt.Printf("My Channel Data3: %d \n", value3) //* Listening to a closed channel will always give result '0'

		wg.Done()
	}(myChannel)

	//* sender
	//* only sender or general channel can close the channel
	go func(channel chan<- int) {
		myChannel <- 7
		//close(myChannel)
		//myChannel <- 7
		//myChannel <- 8
		myChannel <- 8
		//close(myChannel)

		close(myChannel)

		wg.Done()
	}(myChannel)

	wg.Wait()
}
