package components

import (
	"fmt"
	"skn-go-practice/src/pkgs/utils"
)

var wg = utils.GetWg()

func getChannel1(channel chan string) { //* Channel is used to make communication between goroutines
	channel <- "Message!" //* It means "Message!" is passed to channel

	wg.Done()
}

func getChannel2(channel chan string) {
	var message string = <-channel //* It means channel content is passed to message

	fmt.Printf("Message is: %s \n", message)

	wg.Done()
}

func Channel() {
	var channel chan string = make(chan string) //* This is how an empty new channel is created. make() to create a new channel is very very important or the communication won't work

	go getChannel1(channel)

	go getChannel2(channel)

}
