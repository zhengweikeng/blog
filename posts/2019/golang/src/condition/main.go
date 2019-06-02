package main

import (
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	condition := false

	go func() {
		time.Sleep(time.Second)
		c.L.Lock()
		condition = true
		c.Signal()
		c.L.Unlock()
	}()

	c.L.Lock()
	for !condition {
		c.Wait()
	}
	c.L.Unlock()
}
