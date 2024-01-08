package components

import (
	"fmt"
	"skn-go-practice/src/pkgs/utils"
)

var wg = utils.GetWg()

func getChannel0(channel chan string) {
	fmt.Println("This is from Channel Function 0")
	fmt.Printf("This is Channel 0 Data: %v \n", channel)

	wg.Done()
}

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

	go getChannel0(channel)

	go getChannel1(channel)

	go getChannel2(channel)

}

func TestConcurrency() {
	go greeter("Wang So!")
	greeter("SKN!")
}

func greeter(name string) {
	fmt.Printf("You are the best, %s \n", name)
}
