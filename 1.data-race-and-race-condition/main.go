package main

import (
	"fmt"
	"sync"
)

var contadora int
var m sync.Mutex
var wg sync.WaitGroup

func incrementa() {
	defer wg.Done()

	m.Lock()
	contadora++
	m.Unlock()
}

func main() {
	goroutines := 100

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go incrementa()
	}

	wg.Wait()

	fmt.Println(contadora)
}
