package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex

func disOrderData(data interface{}, wg *sync.WaitGroup, i int) {
	fmt.Println(data, i)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	data1 := append(make([]interface{}, 0), "coba 1", "coba 2", "coba 3")
	data2 := append(make([]interface{}, 0), "bisa 1", "bisa 2", "bisa 3")

	wg.Add(8)

	for i := 1; i <= 4; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		go disOrderData(data1, &wg, i)
		go disOrderData(data2, &wg, i)
	}

	wg.Wait()
}
