package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		counter int
		mutex   sync.Mutex
		count   = 1000
		wg      = sync.WaitGroup{}
	)

	wg.Add(count)
	for i := 0; i < count; i += 1 {

		go func(id int) {
			defer wg.Done()

			fmt.Printf("Start gorutine № %v\n", id)

			mutex.Lock()
			counter += 1
			mutex.Unlock()

			fmt.Printf("Stop gorutine № %v\n", id)

		}(i)
	}

	wg.Wait()

	fmt.Printf("\n")
	fmt.Printf("Counter: %v", counter)

}
