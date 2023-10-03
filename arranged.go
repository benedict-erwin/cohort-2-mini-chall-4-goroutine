package main
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var mutex sync.Mutex

func arrangedData(data interface{}, wg *sync.WaitGroup, mutex *sync.Mutex, i int) {
	mutex.Lock()
	fmt.Println(data, i)
	mutex.Unlock()
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	data1 := append(make([]interface{}, 0), "coba 1", "coba 2", "coba 3")
	data2 := append(make([]interface{}, 0), "bisa 1", "bisa 2", "bisa 3")

	wg.Add(8)

	for i := 1; i <= 4; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		go arrangedData(data1, &wg, i)
		go arrangedData(data2, &wg, i)
	}

	wg.Wait()
}
