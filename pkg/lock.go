package pkg

import "sync"

var mutex = &sync.Mutex{}

func Lock() {
	mutex.Lock()
}

func UnLock() {
	mutex.Unlock()
}
