package main

import (
	"fmt"
	"github.com/fztcjjl/ztimer"
	"sync"
	"time"
)

func main() {
	t := ztimer.NewTimer(1, 20)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		time.AfterFunc(time.Duration(i)*time.Second, func() {
			fmt.Println("timer task executed")
			wg.Done()
		})
	}
	t.Start()
	wg.Wait()
}
