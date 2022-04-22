package main

import "fmt"

/*type Product struct{

}
func main(){
	ch:=make(chan Product) //unbuffered channel tanımlamasıdır. Unbuffered channel dediğimiz şey ise
	//sadece bir data alabilen  ve bir data aldığında data başka bir alıcı tarafından okunana kadar bloklanan
	//yapılardır

	ch2:=make(chan Product,10)//buffered channel burada buffered channel da belli size verilir
	//örneğin 10 burada olduğu gibi 10 tane veri gelene kadar bu channel bloklanmaz

	//Channellar go da aslında içerisinde mutex barındıran ve arka planda bir array pointer tutan yapılardır
	//channellar da race condition yaşanmaz ama deadlock yaşanabilir.
}*/

func writeMessage(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Hello, %d", i)
		//üstteki satırdaki <- bu işaret ch'ye (yani channel a) sağdaki datayı yaz demektir
		fmt.Printf("Wrote %d to channel\n", i)
	}
	close(ch) //Eğer burada close fonksiyonunu çağırmazsak okunacak değer bittiğinde
	//ve kod channeldan veri okumaya devam ettiğinde daha veri olmadığı için kod deadlock düşer
}

func main() {
	//ch := make(chan string) //=>unbuffered
	//defer close(ch) => bu yapılmamalı alıcı tarafı channel ı kapatmamalı
	ch := make(chan string, 5) //=>buffered
	go writeMessage(ch)

	for value := range ch {
		fmt.Println(value)
	}
}
