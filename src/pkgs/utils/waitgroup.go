package utils

import "sync"

var wg sync.WaitGroup = sync.WaitGroup{}

var pWg *sync.WaitGroup = &wg

func GetWg() *sync.WaitGroup {

	return pWg
}
