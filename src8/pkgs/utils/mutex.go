package utils

import "sync"

var mutex *sync.Mutex = &sync.Mutex{}
var rwMutex *sync.RWMutex = &sync.RWMutex{}

func GetMutex() *sync.Mutex {
	return mutex
}

func GetRWMutex() *sync.RWMutex {
	return rwMutex
}
