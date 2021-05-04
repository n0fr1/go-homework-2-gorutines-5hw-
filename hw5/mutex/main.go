package main

import (
	"fmt"
	"sync"
)

type Alldata struct {
	result int
	mutex  sync.Mutex
	wg     sync.WaitGroup
}

func main() {

	var (
		getCount Alldata
		num      = 1000
	)

	getCount.wg.Add(num)

	for i := 0; i < num; i += 1 {
		go increment(&getCount)
	}

	getCount.wg.Wait()

	fmt.Printf("Num: %v\n", num)
	fmt.Printf("Result: %v", getCount.result)

}

func increment(a *Alldata) {

	defer a.mutex.Unlock()
	a.mutex.Lock()
	a.result += 1
	a.wg.Done()

}
