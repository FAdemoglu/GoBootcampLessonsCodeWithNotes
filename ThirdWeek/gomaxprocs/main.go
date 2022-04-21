package main

import (
	"fmt"
	"time"
)

func main() {

	/**Gomaxprocs un varsayılan değeri sizin işlemcinizdeki
	çalışan kor kadar
	*/
	//runtime.GOMAXPROCS(1) bu önerilen bir yöntem değil bu bir sadece bir core kullanıcağını belirtmek için
	//Maximum logical cpu sayısı değiştirme
	go func(i int) {
		for letter := 'c'; letter < 'c'+int32(i); letter++ {
			fmt.Printf("%c ", letter)
		}
	}(250)

	go func(i int) {
		for j := 0; j < i; j++ {
			fmt.Printf("%d ", j)
			//time.Sleep(time.Millisecond * 5)
		}
	}(300)

	time.Sleep(time.Second * 10)
	/*fmt.Println("Hello world")

		go func() {
			fmt.Println("Hello from goroutine")
		}()
	//Burada time sleep i koymazsak go func içindeki kısım denk gelirse çaışır
	//Sleep i koyduğumuz zaman çalışabilir fakat bu yine garanti etmez sebebi ise şu
	//main main goroutine olduğu için main goroutine sonlandığında içerde çalışan tüm
	//goroutinelerde exit olur

		time.Sleep(time.Second * 2)*/
}
