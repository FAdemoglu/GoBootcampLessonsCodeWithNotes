package main

import (
	"fmt"
	"sync"
)

//Sync once sadece bir defa çalıştırmak istediğiniz fonksiyonlar için kullanılabilen bir methoddur concurrent safedir

//Sync once nerelerde kullanılabilir sadece bir struct yaratılıcaksa, bir işlem yapılacak ve sadece bir kere yapılacaksa
func main() {
	//SyncOnce()
	//SyncOnceDeadlock()
	SyncOnceWithMultipleFunction()
}

func SyncOnce() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
			increment()
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

func SyncOnceWithMultipleFunction() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	//Burada count değeri 1 olur sebebi ise once.Do sadece bir kere çalışır yani burada increment i çalıştırır
	//ve sonuç sadece 1 olur

	fmt.Printf("Count: %d\n", count)
}

func SyncOnceDeadlock() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }

	onceA.Do(initA)
}
