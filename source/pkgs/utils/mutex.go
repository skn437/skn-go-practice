package utils

import "sync"

var mutex *sync.Mutex = &sync.Mutex{}

func GetMutex() *sync.Mutex {
	return mutex
}
