package main

import (
	"fmt"
	"time"
)

/**
Select channellarda okuma ve yazam işlemlerinde birden fazla channel ı dinleyebilmek için
oluşturulmuş bir kontrol yapısıdır yazımı swtich case e benzer fakat switch case değildir farklı bir çalışma
yapısı vardır
*/

func getCustomerId(ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- 20
}

func main() {
	ch := make(chan int)

	go getCustomerId(ch)

	for {
		time.Sleep(1 * time.Second)
		select {
		///Burada iki casse durumunda eğer iki channela da data gelirse random birisine girer
		case a := <-ch:
			fmt.Println(a)
			return
		default:
			fmt.Println("No data")
		}

	}
}
