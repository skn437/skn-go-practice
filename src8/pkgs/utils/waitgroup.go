package utils

import "sync"

var wg *sync.WaitGroup = &sync.WaitGroup{}

func GetWaitGroup() *sync.WaitGroup {
	return wg
}
