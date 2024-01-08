package utils

import "sync"

var mutex *sync.Mutex = &sync.Mutex{}

func GetMutex() *sync.Mutex {
	return mutex
}

var mutexRW *sync.RWMutex = &sync.RWMutex{}

func GetRWMutex() *sync.RWMutex {
	return mutexRW
}
