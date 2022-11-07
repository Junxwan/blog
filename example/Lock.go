package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 0
	num := 3
	wait := new(sync.WaitGroup)
	wait.Add(num)
	lock := new(sync.Mutex)

	for i := 0; i < num; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				lock.Lock()
				a += 1
				lock.Unlock()
			}

			wait.Done()
		}()
	}

	wait.Wait()

	fmt.Print(a) // 30000


}
