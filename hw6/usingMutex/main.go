//1. Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких
//потоков. Выполните трассировку программы.

package main

import (
	"fmt"
	"sync"
)

func main() {

	var mutex sync.Mutex
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {

		defer wg.Done()

		for c := 'a'; c <= 'z'; c += 1 {
			mutex.Lock()
			fmt.Printf("%c", c)
			mutex.Unlock()
		}

	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 20; i += 1 {
			mutex.Lock()
			fmt.Printf("%d", i)
			mutex.Unlock()
		}

	}()

	wg.Wait()

}
