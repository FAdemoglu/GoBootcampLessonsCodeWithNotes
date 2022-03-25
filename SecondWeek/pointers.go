package main

import "fmt"

//Pointer paylaşılabilen ama kopyalanmayan değerlerdir value lar ise
//kopyalanabilen ama paylaşılam değerleri ifade eder.

func main() {
	a := Album{"Album Name", 10}

	fmt.Printf("Album sold :%d", a.Sold)
	//Değer kopyalanarak gönderiliyor bu sebeple sold değeri güncellenmiyor
	//Değer kopyalanarak çapıral memthod
	//sell(a, 10)
	sell(&a, 10)
	fmt.Printf("[Main] After sell function Album sold : %d", a.Sold)
}

func sell(a *Album, qty int) {
	a.Sold += qty
	fmt.Printf("[sell] Inside sell function Album sold : %d", a.Sold)
}

//Değer kopyalanarak gönderilen fonksiyon
/*func sell(a Album, qty int) {
	a.Sold += qty
	fmt.Printf("[sell] Inside sell function Album sold : %d", a.Sold)
}*/

type Album struct {
	Name string
	Sold int
}

// Çoğu zaman 64byte lık kısımlardan fazla alanlar için pointerla gönderip alım işlemi yapılmalı

//Pointerlar method içerisinde bir değeri değiştirmek istiyorsanız kullanılır
