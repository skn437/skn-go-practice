package components

import (
	"fmt"
	"net/http"
	"skn-go-practice/source/pkgs/utils"
	"sync"
)

var wg *sync.WaitGroup = utils.GetWaitGroup()

var mutex *sync.Mutex = utils.GetMutex()

var websites []string = []string{
	"https://google.com",
	"https://youtube.com",
	"https://go.dev",
	"https://github.com",
}

var Signals *[]string = &[]string{}

func GetWebsiteList() *[]string {
	return &websites
}

func GetStatusCode(endpoint string) {
	response, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	mutex.Lock()

	*Signals = append(*Signals, endpoint)

	mutex.Unlock()

	fmt.Printf("%d status code for endpoint %s \n", response.StatusCode, endpoint)

	wg.Done()
}
