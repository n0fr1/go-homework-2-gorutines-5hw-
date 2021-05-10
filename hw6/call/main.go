//2. Написать многопоточную программу, в которой будет использоваться явный вызов
//планировщика. Выполните трассировку программы
//трассировка:
//1. go run main.go 2>trace.out - перенаправляем вывод в trace.out
//2. go tool trace trace.out    - трассировка в графическом виде

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {

	counter := 0

	const gs = 100
	var wg sync.WaitGroup

	trace.Start(os.Stderr)
	defer trace.Stop()

	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()

			counter++
			if counter == 10 {
				runtime.Gosched()
			}

			fmt.Println("count:", counter)

			mu.Unlock()
			wg.Done()
		}()
	}

	fmt.Println("test")

	wg.Wait()

}
