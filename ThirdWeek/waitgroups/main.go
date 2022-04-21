package main

import (
	"fmt"
	"sync"
)

func main() {

	//Burada rastgele sayılar yazar ve bunun ana sıkıntısının sebebi şu kendisini kapsayan bir fonksiyondaki
	//değişkene ulaşıyor olması
	//Bunun çözümü için parametre olarak geçilebilir. Ama burada parametre olarak geçtikten sonra
	//değerler hala sıralı bir biçimde ekrana basılmaz bunun içinde wait group denen bir yapı karşımıza çıkar
	wg := sync.WaitGroup{} //Wait group concurrent safe bir counter
	//Bir üstteki satırda boş bir waitgroup yarattık
	//wg.Add(10)//Burada 10 tane goroutine yarattık ve bu 10 tane goroutine beklenilecek burada deadlock oluşur
	for i := 0; i < 10; i++ {
		wg.Add(1) //burada tek tekte verilebilir
		go func(value int) {
			fmt.Println(value)
			wg.Done() //bu done fonksiyonu yukarıda verilen 10 değerini bir azaltıyor.
		}(i)
		wg.Wait() //Burada wait yukarıda verilen 10 değerini sıfıra ulaşana kadar beklemesi için
	}

	//time.Sleep(time.Second*2)
}
