package main

import (
	"fmt"
	"math"
)

/**
Go da aslında her şey bir packagedır

Bütün go dosyaları bir paket tanımlamasıyla başlar

Main fonksiyonu bütün fonksiyonlar için bir entry pointtir.

Go da bir değişken tanımlanıyorsa kullanılmak zorunda aynı şey paketler ve kütüphanler için de geçerli

Go da fonksiyonlar da birden fazla değer dönülebilir
*/

//salary:=10 => Bu tanımlama paket seviyesinde bir tanımlama ayrıca bu short declaration burada tanımlanamaz paket seviyesinde hata verir

//Go dilinde bir erişim belirleyicisi yok bir başka paketteki fonksiyona vs erişmek için
// o fonksiyonun adının baş harfi büyük olmalı bu public yapar küçük harf private yapar

var GlobalIntVar = 100

func main() {
	//Go da 3 şekilde değişken ataması yapılabiliyor.
	//var myVar int =64
	//var myName = "patika"
	//salary:=10 => Short declaration Bu değişkenler fonksiyon seviyesinde yapılan bir tanımlama
	fmt.Printf("Hello World")

	sum := GlobalIntVar - 50 //Global değişken bu şekilde kullanılıyor

	fmt.Printf("Total value : %d", sum) //Integer i string format lama ve ekranda bastırma

	sumAbs := math.Abs(float64(sum))

	fmt.Printf("Sum Abs result : %v", sumAbs)

}
