//2. Написать многопоточную программу, в которой будет использоваться явный вызов
//планировщика. Выполните трассировку программы
//трассировка:
//1. go run main.go 2>trace.out - перенаправляем вывод в trace.out
//2. go tool trace trace.out    - трассировка в графическом виде

package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {

	trace.Start(os.Stderr)
	defer trace.Stop()

	go log.Println("I'm working!")
	for i := 0; ; i += 1 {
		if i%1e6 == 0 {
			runtime.Gosched() //start sheduler
		}
	}
}
