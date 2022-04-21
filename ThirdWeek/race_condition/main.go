package main

import (
	"fmt"
	"sync"
)

//race condition=>paylaşılan bir değerin birden fazla goroutine tarafından update edilmesiyle sebebiyle yanlış sonuçlar elde etme durumudur.
//Bunu önlemenin bir yolu mutexlerdir bu tarz paylaşılan erişim noktalarını kontrol etme imkanı sağlaya bir yapıdır

var counter = 0

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)
	mut := sync.Mutex{}
	rwmut := sync.RWMutex{}
	go func() {
		for i := 0; i < 10000; i++ {
			mut.Lock()
			counter++
			mut.Unlock()
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mut.Lock()
			counter++
			mut.Unlock()

		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(counter)
}
