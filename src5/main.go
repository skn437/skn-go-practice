package main

import (
	"fmt"
	"skn-go-practice/src5/pkgs/utils"
	"time"
)

// * Go `time` stdlib practice examples
func main() {
	var timeNow time.Time = time.Now()

	fmt.Printf("Time Right Now: %v \n", timeNow)
	fmt.Printf("Time Right Now Formatted: %v \n", timeNow.Format("2006-01-02 15:04:05 Monday")) //* The format is: `yyyy-mm-dd`. You must put 'Monday' in order to get the day. You must use '15:04:05' in order to get the time in `hh:mm:ss` format.

	var newTime time.Time = time.Date(2022, time.August, 9, 17, 30, 0, 0, time.Local) //* New Date Creation

	fmt.Printf("New Time: %v \n", newTime.Format("2006-01-02 15:04:05 Monday"))

	var user utils.UserType = utils.UserType{
		Name:  "SKN",
		Count: 7,
	}

	fmt.Println(user.Count)

}
