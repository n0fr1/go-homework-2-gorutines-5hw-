//3.Смоделировать ситуацию “гонки”, и проверить программу на наличии “гонки”.
//проверка на состояние гонки: go run -race main.go

package main

import (
	"fmt"
	"sync"
)

const count = 1000

func main() {
	var (
		counter int
		wg      sync.WaitGroup
	)
	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func() {
			defer wg.Done()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

/*
==================
WARNING: DATA RACE
Read at 0x00c0000b4010 by goroutine 8:
  main.main.func1()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:22 +0x6c

Previous write at 0x00c0000b4010 by goroutine 7:
  main.main.func1()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:22 +0x82

Goroutine 8 (running) created at:
  main.main()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:20 +0xe4

Goroutine 7 (finished) created at:
  main.main()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:20 +0xe4
==================
==================
WARNING: DATA RACE
Read at 0x00c0000b4010 by goroutine 9:
  main.main.func1()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:22 +0x6c

Previous write at 0x00c0000b4010 by goroutine 7:
  main.main.func1()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:22 +0x82

Goroutine 9 (running) created at:
  main.main()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:20 +0xe4

Goroutine 7 (finished) created at:
  main.main()
      /home/evgeniy/go-homework-2-gorutines-5hw-/hw6/race/main.go:20 +0xe4
==================
979
Found 2 data race(s)
exit status 66

*/
