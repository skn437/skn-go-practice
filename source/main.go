package main

import (
	"fmt"
	"skn-go-practice/source/pkgs/components"
	"skn-go-practice/source/pkgs/utils"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = utils.GetWaitGroup()

	var websites *[]string = components.GetWebsiteList()

	wg.Add(len(*websites))

	for _, website := range *websites {
		go components.GetStatusCode(website)
	}

	wg.Wait()

	fmt.Println(*components.Signals)
}
