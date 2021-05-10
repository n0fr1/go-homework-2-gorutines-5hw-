//1. Написать программу, которая использует мьютекс для безопасного доступа к данным из нескольких
//потоков. Выполните трассировку программы.
//трассировка:
//1. go run main.go 2>trace.out - перенаправляем вывод в trace.out
//2. go tool trace trace.out    - трассировка в графическом виде

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {

	var mutex sync.Mutex
	wg := sync.WaitGroup{}

	trace.Start(os.Stderr)
	defer trace.Stop()

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
